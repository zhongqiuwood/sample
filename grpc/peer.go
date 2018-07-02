package p2p

import (
	"golang.org/x/net/context"

	"github.com/op/go-logging"
	"github.com/golang/protobuf/proto"
	"fmt"
	"io"
	"sync"
	"google.golang.org/grpc"
	"github.com/zhongqiuwood/sample/grpc/comm"
	pb "github.com/zhongqiuwood/sample/grpc/protos"

	"github.com/golang/protobuf/ptypes/timestamp"
	"time"
)

var peerLogger = logging.MustGetLogger("peer")


type handlerMap struct {
	sync.RWMutex
	m map[string]*PeerHandler
}



type PeerImpl struct {
	//handlerFactory HandlerFactory
	handlerMap     *handlerMap
	endpoint       *pb.PeerEndpoint
	//secHelper      crypto.Peer
	//engine         Engine
	//isValidator    bool
	//reconnectOnce  sync.Once
	//discHelper     discovery.Discovery
	//discPersist    bool
	//mode 		   string
	//trustedValidators []*pb.PeerID
}

func NewPeer(id string, addr string) (peer *PeerImpl, err error) {
	peer = new(PeerImpl)

	peer.handlerMap = &handlerMap{m: make(map[string]*PeerHandler)}

	peer.endpoint = &pb.PeerEndpoint{id, addr, pb.PeerEndpoint_UNDEFINED}
	return peer, nil

}

// ChatStream interface supported by stream between Peers
type ChatStream interface {
	Send(message *pb.OkMessage) error
	Recv() (*pb.OkMessage, error)
}

func NewPeerClientConnectionWithAddress(peerAddress string) (*grpc.ClientConn, error) {

	return comm.NewClientConnectionWithAddress(peerAddress, true, false, nil)
}


func (p *PeerImpl) ChatWithSomePeers(addresses []string) {

	if len(addresses) == 0 {
		peerLogger.Debug("Starting up the first peer of a new network")
		return // nothing to do
	}
	for _, address := range addresses {

		peerLogger.Infof("Try to connect to: %s", address)
		go p.chatWithPeer(address)
	}
}

func (p *PeerImpl) chatWithPeer(address string) error {
	peerLogger.Debugf("Initiating Chat with peer address: %s", address)
	conn, err := NewPeerClientConnectionWithAddress(address)
	if err != nil {
		peerLogger.Errorf("Error creating connection to peer address %s: %s", address, err)
		return err
	}
	serverClient := pb.NewPeerClient(conn)
	ctx := context.Background()
	stream, err := serverClient.Chat(ctx)
	if err != nil {
		peerLogger.Errorf("Error establishing chat with peer address %s: %s", address, err)
		return err
	}
	peerLogger.Debugf("Established Chat with peer address: %s", address)
	err = p.handleChat(ctx, stream, true)
	stream.CloseSend()
	if err != nil {
		peerLogger.Errorf("Ending Chat with peer address %s due to error: %s", address, err)
		return err
	}
	return nil
}

// Chat implementation of the the Chat bidi streaming RPC function
func (p *PeerImpl) handleChat(ctx context.Context, stream ChatStream, initiatedStream bool) error {
	deadline, ok := ctx.Deadline()
	peerLogger.Debugf("Current context deadline = %s, ok = %v, initiatedStream=%t", deadline, ok, initiatedStream)
	//handler, err := p.handlerFactory(p, stream, initiatedStream)
	handler, err := NewPeerHandler(stream, initiatedStream, p)

	if err != nil {
		return fmt.Errorf("Error creating handler during handleChat initiation: %s", err)
	}
	defer handler.Stop()
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			peerLogger.Debug("Received EOF, ending Chat")
			return nil
		}
		if err != nil {
			e := fmt.Errorf("Error during Chat, stopping handler: %s", err)
			peerLogger.Error(e.Error())
			return e
		}
		err = handler.HandleMessage(in)
		if err != nil {
			if in.GetType() != pb.OkMessage_DISC_GET_PEERS &&
				in.GetType() != pb.OkMessage_DISC_HELLO {
				peerLogger.Errorf("Error handling message: %s", err)
			}
			return err
		}
	}
	return nil
}



func getHandlerKey(peerMessageHandler *PeerHandler) (string, error) {
	peerEndpoint, err := peerMessageHandler.To()
	if err != nil {
		return "", fmt.Errorf("Error getting messageHandler key: %s", err)
	}
	return peerEndpoint.ID, nil
}



func (p *PeerImpl) newHelloMessage(msg string) (*pb.HelloMessage, error) {

	peerLogger.Debugf("newHelloMessage: %+v", p.endpoint)
	return &pb.HelloMessage{p.endpoint}, nil
}

func (p *PeerImpl) NewOpenchainDiscoveryHello(msg string) (*pb.OkMessage, error) {
	helloMessage, err := p.newHelloMessage(msg)
	if err != nil {
		return nil, fmt.Errorf("???Error getting new HelloMessage: %s", err)
	}
	data, err := proto.Marshal(helloMessage)
	if err != nil {
		return nil, fmt.Errorf("Error marshalling HelloMessage: %s", err)
	}
	// Need to sign the Discovery Hello message
	newDiscoveryHelloMsg := &pb.OkMessage{Type: pb.OkMessage_DISC_HELLO, Payload: data, Timestamp: CreateUtcTimestamp()}
	//err = p.signMessageMutating(newDiscoveryHelloMsg)
	//if err != nil {
	//	return nil, fmt.Errorf("Error signing new HelloMessage: %s", err)
	//}
	return newDiscoveryHelloMsg, nil
}

// RegisterHandler register a MessageHandler with this coordinator
func (p *PeerImpl) RegisterHandler(messageHandler *PeerHandler) error {
	key, err := getHandlerKey(messageHandler)
	if err != nil {
		return fmt.Errorf("Error registering handler: %s", err)
	}
	p.handlerMap.Lock()
	defer p.handlerMap.Unlock()
	if _, ok := p.handlerMap.m[key]; ok == true {
		// Duplicate, return error
		//return newDuplicateHandlerError(messageHandler)
		return fmt.Errorf("Error creating Duplicate Handler error: %+v", messageHandler)
	}

	p.handlerMap.m[key] = messageHandler
	peerLogger.Debugf("registered handler with key: %s", key)
	return nil
}

// DeregisterHandler deregisters an already registered MessageHandler for this coordinator
func (p *PeerImpl) DeregisterHandler(messageHandler *PeerHandler) error {
	key, err := getHandlerKey(messageHandler)
	if err != nil {
		return fmt.Errorf("Error deregistering handler: %s", err)
	}
	p.handlerMap.Lock()
	defer p.handlerMap.Unlock()
	if _, ok := p.handlerMap.m[key]; !ok {
		// Handler NOT found
		return fmt.Errorf("Error deregistering handler, could not find handler with key: %s", key)
	}
	delete(p.handlerMap.m, key)
	peerLogger.Debugf("Deregistered handler with key: %s", key)
	return nil
}

// RPC RouteChat receives a stream of message/location pairs, and responds with a stream of all
// previous messages at each of those locations.
func (p *PeerImpl) Chat(stream pb.Peer_ChatServer) error {
	return p.handleChat(stream.Context(), stream, false)
}


// CreateUtcTimestamp returns a google/protobuf/Timestamp in UTC
func CreateUtcTimestamp() *timestamp.Timestamp {
	now := time.Now().UTC()
	secs := now.Unix()
	nanos := int32(now.UnixNano() - (secs * 1000000000))
	return &(timestamp.Timestamp{Seconds: secs, Nanos: nanos})
}