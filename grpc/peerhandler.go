package p2p

import (
	"sync"
	"time"
	"fmt"
	"github.com/looplab/fsm"
	"github.com/spf13/viper"
	"github.com/golang/protobuf/proto"

	pb "github.com/zhongqiuwood/sample/grpc/protos"

)

const DefaultSyncSnapshotTimeout time.Duration = 60 * time.Second

// Handler peer handler implementation.
type PeerHandler struct {
	chatMutex                     sync.Mutex
	ChatStream                    ChatStream
	Coordinator                   *PeerImpl
	doneChan                      chan struct{}
	FSM                           *fsm.FSM
	initiatedStream               bool // Was the stream initiated within this Peer
	registered                    bool
	syncSnapshotTimeout           time.Duration
	lastIgnoredSnapshotCID        *uint64
	ToPeerEndpoint                *pb.PeerEndpoint
}

func NewPeerHandler(stream ChatStream, initiatedStream bool, peer *PeerImpl) (*PeerHandler, error) {

	d := &PeerHandler{
		ChatStream:      stream,
		initiatedStream: initiatedStream,
		Coordinator:peer,
	}
	d.doneChan = make(chan struct{})

	if dur := viper.GetDuration("peer.sync.state.snapshot.writeTimeout"); dur == 0 {
		d.syncSnapshotTimeout = DefaultSyncSnapshotTimeout
	} else {
		d.syncSnapshotTimeout = dur
	}

	d.FSM = fsm.NewFSM(
		"created",
		fsm.Events{
			{Name: pb.OkMessage_DISC_HELLO.String(), Src: []string{"created"}, Dst: "established"},
			//{Name: pb.Message_DISC_GET_PEERS.String(), Src: []string{"established"}, Dst: "established"},
			//{Name: pb.Message_DISC_PEERS.String(), Src: []string{"established"}, Dst: "established"},
			//{Name: pb.Message_SYNC_BLOCK_ADDED.String(), Src: []string{"established"}, Dst: "established"},
			//{Name: pb.Message_SYNC_GET_BLOCKS.String(), Src: []string{"established"}, Dst: "established"},
			//{Name: pb.Message_SYNC_GET_BLOCKS_RESP.String(), Src: []string{"established"}, Dst: "established"},
			//{Name: pb.Message_SYNC_STATE_GET_SNAPSHOT.String(), Src: []string{"established"}, Dst: "established"},
			//{Name: pb.Message_SYNC_STATE_SNAPSHOT.String(), Src: []string{"established"}, Dst: "established"},
			//{Name: pb.Message_SYNC_STATE_GET_DELTAS.String(), Src: []string{"established"}, Dst: "established"},
			//{Name: pb.Message_SYNC_STATE_DELTAS.String(), Src: []string{"established"}, Dst: "established"},
			//{Name: pb.Message_SYNC_GET_STATE_HASH.String(), Src: []string{"established"}, Dst: "established"},
			//{Name: pb.Message_SYNC_GET_STATE_HASH_RESP.String(), Src: []string{"established"}, Dst: "established"},
		},
		fsm.Callbacks{
			"enter_state":                                           func(e *fsm.Event) { d.enterState(e) },
			"before_" + pb.OkMessage_DISC_HELLO.String():              func(e *fsm.Event) { d.beforeHello(e) },
			//"before_" + pb.Message_DISC_GET_PEERS.String():          func(e *fsm.Event) { d.beforeGetPeers(e) },
			//"before_" + pb.Message_DISC_PEERS.String():              func(e *fsm.Event) { d.beforePeers(e) },
			//"before_" + pb.Message_SYNC_BLOCK_ADDED.String():        func(e *fsm.Event) { d.beforeBlockAdded(e) },
			//"before_" + pb.Message_SYNC_GET_BLOCKS.String():         func(e *fsm.Event) { d.beforeSyncGetBlocks(e) },
			//"before_" + pb.Message_SYNC_GET_BLOCKS_RESP.String():    func(e *fsm.Event) { d.beforeSyncBlocks(e) },
			//"before_" + pb.Message_SYNC_STATE_GET_SNAPSHOT.String(): func(e *fsm.Event) { d.beforeSyncStateGetSnapshot(e) },
			//"before_" + pb.Message_SYNC_STATE_SNAPSHOT.String():     func(e *fsm.Event) { d.beforeSyncStateSnapshot(e) },
			//"before_" + pb.Message_SYNC_STATE_GET_DELTAS.String():   func(e *fsm.Event) { d.beforeSyncStateGetDeltas(e) },
			//"before_" + pb.Message_SYNC_STATE_DELTAS.String():       func(e *fsm.Event) { d.beforeSyncStateDeltas(e) },
			//"before_" + pb.Message_SYNC_GET_STATE_HASH.String():         func(e *fsm.Event) { d.beforeGetStateHash(e) },
			//"before_" + pb.Message_SYNC_GET_STATE_HASH_RESP.String():    func(e *fsm.Event) { d.beforeGetStateHashResp(e) },
		},
	)

	// If the stream was initiated from this Peer, send an Initial HELLO message
	if d.initiatedStream {
		// Send intiial Hello
		helloMessage, err := d.Coordinator.NewOpenchainDiscoveryHello("this is the client")
		if err != nil {
			return nil, fmt.Errorf("??????Error getting new HelloMessage: %s", err)
		}
		if err := d.SendMessage(helloMessage); err != nil {
			return nil, fmt.Errorf("Error creating new Peer Handler, error returned sending %s: %s", pb.OkMessage_DISC_HELLO, err)
		}
	}

	return d, nil
}


// HandleMessage handles the Openchain messages for the Peer.
func (d *PeerHandler) HandleMessage(msg *pb.OkMessage) error {

	//if d.ToPeerEndpoint != nil && d.ToPeerEndpoint.ID != nil {
	//	if msg.Type != pb.Message_DISC_PEERS && msg.Type != pb.Message_DISC_GET_PEERS {
	//		debugger.Log(debugger.INFO,"<<<--- Peer handler received message<%s> from <%s>",
	//			msg.Type.String(), d.ToPeerEndpoint.ID)
	//	}
	//}

	if d.FSM.Cannot(msg.Type.String()) {
		//debugger.Log(debugger.ERROR,"Peer<%s> FSM cannot handle message (%s) with size (%d) while in state: %s",
		//	d.ToPeerEndpoint.ID, msg.Type.String(), len(msg.Payload), d.FSM.Current())
		return fmt.Errorf("Peer FSM cannot handle message (%s) with payload size (%d) while in state: %s",
			msg.Type.String(), len(msg.Payload), d.FSM.Current())
	}
	err := d.FSM.Event(msg.Type.String(), msg)
	if err != nil {
		if _, ok := err.(*fsm.NoTransitionError); !ok {
			// Only allow NoTransitionError's, all others are considered true error.
			return fmt.Errorf("Peer FSM failed while handling message (%s): current state: %s, error: %s",
				msg.Type.String(), d.FSM.Current(), err)
			//t.Error("expected only 'NoTransitionError'")
		}
	}
	return nil
}

// SendMessage sends a message to the remote PEER through the stream
func (d *PeerHandler) SendMessage(msg *pb.OkMessage) error {
	//make sure Sends are serialized. Also make sure everyone uses SendMessage
	//instead of calling Send directly on the grpc stream
	d.chatMutex.Lock()
	defer d.chatMutex.Unlock()

	//if msg.Type != pb.Message_DISC_PEERS &&
	//	msg.Type != pb.Message_DISC_GET_PEERS &&
	//	msg.Type != pb.Message_DISC_HELLO  {
	//	if d.ToPeerEndpoint != nil && d.ToPeerEndpoint.ID != nil {
	//		debugger.Log(debugger.INFO,"<<<--- Sending msg<%s> payload<%s> to peer <%s>",
	//			msg.Type.String(),	pb.Consensus_Payload_Type(msg.PayloadType).String(), d.ToPeerEndpoint.ID)
	//	}
	//}

	err := d.ChatStream.Send(msg)
	if err != nil {
		return fmt.Errorf("Error Sending message through ChatStream: %s", err)
	}
	return nil
}

func (d *PeerHandler) Stop() error {
	// Deregister the handler
	err := d.deregister()
	if err != nil {
		return fmt.Errorf("Error stopping MessageHandler: %s", err)
	}
	return nil
}

func (d *PeerHandler) deregister() error {
	var err error
	if d.registered {
		err = d.Coordinator.DeregisterHandler(d)
		//doneChan is created and waiting for registered handlers only
		d.doneChan <- struct{}{}
		d.registered = false
	}
	return err
}

func (d *PeerHandler) enterState(e *fsm.Event) {
	peerLogger.Debugf("The Peer's bi-directional stream to %s is %s, from event %s\n", d.ToPeerEndpoint, e.Dst, e.Event)
	//debugger.Log(debugger.INFO,"Bi-directional to <%s> is %s, from event %s\n", d.ToPeerEndpoint, e.Dst, e.Event)
}


func (d *PeerHandler) beforeHello(e *fsm.Event) {
	peerLogger.Debugf("Received %s, parsing out Peer identification", e.Event)
	// Parse out the PeerEndpoint information
	if _, ok := e.Args[0].(*pb.OkMessage); !ok {
		e.Cancel(fmt.Errorf("Received unexpected message type"))
		return
	}
	msg := e.Args[0].(*pb.OkMessage)

	helloMessage := &pb.HelloMessage{}
	err := proto.Unmarshal(msg.Payload, helloMessage)
	if err != nil {
		e.Cancel(fmt.Errorf("Error unmarshalling HelloMessage: %s", err))
		return
	}
	// Store the PeerEndpoint
	d.ToPeerEndpoint = helloMessage.Endpoint
	peerLogger.Debugf("Received %s from endpoint=%s", e.Event, helloMessage)

	if d.initiatedStream == false {
		// Did NOT intitiate the stream, need to send back HELLO
		peerLogger.Debugf("Received %s, sending back %s", e.Event, pb.OkMessage_DISC_HELLO.String())
		// Send back out PeerID information in a Hello
		helloMessage, err := d.Coordinator.NewOpenchainDiscoveryHello("this is the server")
		if err != nil {
			e.Cancel(fmt.Errorf("??Error getting new HelloMessage: %s", err))
			return
		}
		if err := d.SendMessage(helloMessage); err != nil {
			e.Cancel(fmt.Errorf("Error sending response to %s:  %s", e.Event, err))
			return
		}
	}
	// Register
	err = d.Coordinator.RegisterHandler(d)
	if err != nil {
		e.Cancel(fmt.Errorf("Error registering Handler: %s", err))
	} else {
		// Registered successfully
		d.registered = true
		//otherPeer := d.ToPeerEndpoint.Address

		//StoreDiscoveryList
		//if !d.Coordinator.GetDiscHelper().FindNode(otherPeer) {
		//	if ok := d.Coordinator.GetDiscHelper().AddNode(otherPeer); !ok {
		//		peerLogger.Warningf("Unable to add peer %v to discovery list", otherPeer)
		//	}
		//	err = d.Coordinator.StoreDiscoveryList()
		//	if err != nil {
		//		peerLogger.Error(err)
		//	}
		//}
		//go d.start()
	}
}

func (d *PeerHandler) To() (pb.PeerEndpoint, error) {
	if d.ToPeerEndpoint == nil {
		return pb.PeerEndpoint{}, fmt.Errorf("No peer endpoint for handler")
	}
	return *(d.ToPeerEndpoint), nil
}

func (d *PeerHandler) start() error {
	discPeriod := viper.GetDuration("peer.discovery.period")
	tickChan := time.NewTicker(discPeriod).C
	peerLogger.Debug("Starting Peer discovery service")
	for {
		select {
		case <-tickChan:
			if err := d.SendMessage(&pb.OkMessage{Type: pb.OkMessage_DISC_GET_PEERS}); err != nil {
				peerLogger.Errorf("Error sending %s during handler discovery tick: %s",
					pb.OkMessage_DISC_GET_PEERS, err)
			}
		case <-d.doneChan:
			peerLogger.Debug("Stopping discovery service")
			return nil
		}
	}
}
