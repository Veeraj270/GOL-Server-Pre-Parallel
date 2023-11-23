package stubs

var RemoteCall string = "RemoteProcessor.CallRemoteDistributor"
var RemoteCellCount string = "RemoteProcessor.CallNumberOfAliveCells"

var RemotePause string = "RemoteProcessor.CallPause"

var RemoteSave string = "RemoteProcessor.CallSave"
var RemoteClose string = "RemoteProcessor.CallClose"

type Request struct {
	World   [][]uint8
	Turns   int
	Threads int
}

type Response struct {
	World [][]uint8
}

type CellCountRequest struct {
}

type CellCountResponse struct {
	Turn      int
	CellCount int
}

type PauseReq struct {
	Paused bool
}

type PauseResp struct {
	Turn int
}

type SaveReq struct {
}

type SaveResp struct {
	World [][]uint8
	Turn  int
}

type CloseReq struct{}

type CloseResp struct{}
