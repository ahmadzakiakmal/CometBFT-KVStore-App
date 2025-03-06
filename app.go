package main

import (
	"context"

	abcitypes "github.com/cometbft/cometbft/abci/types"
)

type KVStoreApplication struct {
}

var _ abcitypes.Application = (*KVStoreApplication)(nil)

func NewKVStoreApplicatio() *KVStoreApplication {
	return &KVStoreApplication{}
}

func (app *KVStoreApplication) Info(
	_ context.Context,
	info *abcitypes.InfoRequest,
) (*abcitypes.InfoResponse, error) {
	return &abcitypes.InfoResponse{}, nil
}

func (app *KVStoreApplication) Query(
	_ context.Context,
	req *abcitypes.QueryRequest,
) (*abcitypes.QueryResponse, error) {
	return &abcitypes.QueryResponse{}, nil
}

func (app *KVStoreApplication) CheckTx(
	_ context.Context,
	check *abcitypes.CheckTxRequest,
) (*abcitypes.CheckTxResponse, error) {
	return &abcitypes.CheckTxResponse{}, nil
}

func (app *KVStoreApplication) InitChain(
	_ context.Context,
	chain *abcitypes.InitChainRequest,
) (*abcitypes.InitChainResponse, error) {
	return &abcitypes.InitChainResponse{}, nil
}

func (app *KVStoreApplication) PrepareProposal(
	_ context.Context,
	proposal *abcitypes.PrepareProposalRequest,
) (*abcitypes.PrepareProposalResponse, error) {
	return &abcitypes.PrepareProposalResponse{}, nil
}

func (app *KVStoreApplication) ProcessProposal(
	_ context.Context,
	proposal *abcitypes.ProcessProposalRequest,
) (*abcitypes.ProcessProposalResponse, error) {
	return &abcitypes.ProcessProposalResponse{}, nil
}

func (app *KVStoreApplication) FinalizeBlock(
	_ context.Context,
	req *abcitypes.FinalizeBlockRequest,
) (*abcitypes.FinalizeBlockResponse, error) {
	return &abcitypes.FinalizeBlockResponse{}, nil
}

func (app *KVStoreApplication) Commit(
	_ context.Context,
	commit *abcitypes.CommitRequest,
) (*abcitypes.CommitResponse, error) {
	return &abcitypes.CommitResponse{}, nil
}

func (app *KVStoreApplication) ListSnapshots(
	_ context.Context,
	snapshots *abcitypes.ListSnapshotsRequest,
) (*abcitypes.ListSnapshotsResponse, error) {
	return &abcitypes.ListSnapshotsResponse{}, nil
}

func (app *KVStoreApplication) OfferSnapshot(
	_ context.Context,
	snapshot *abcitypes.OfferSnapshotRequest,
) (*abcitypes.OfferSnapshotResponse, error) {
	return &abcitypes.OfferSnapshotResponse{}, nil
}

func (app *KVStoreApplication) LoadSnapshotChunk(
	_ context.Context,
	chunk *abcitypes.LoadSnapshotChunkRequest,
) (*abcitypes.LoadSnapshotChunkResponse, error) {
	return &abcitypes.LoadSnapshotChunkResponse{}, nil
}

func (app *KVStoreApplication) ApplySnapshotChunk(
	_ context.Context,
	chunk *abcitypes.ApplySnapshotChunkRequest) (*abcitypes.ApplySnapshotChunkResponse, error) {
	return &abcitypes.ApplySnapshotChunkResponse{
		Result: abcitypes.APPLY_SNAPSHOT_CHUNK_RESULT_ACCEPT,
	}, nil
}

func (app *KVStoreApplication) ExtendVote(
	_ context.Context,
	extend *abcitypes.ExtendVoteRequest,
) (*abcitypes.ExtendVoteResponse, error) {
	return &abcitypes.ExtendVoteResponse{}, nil
}

func (app *KVStoreApplication) VerifyVoteExtension(
	_ context.Context,
	verify *abcitypes.VerifyVoteExtensionRequest,
) (*abcitypes.VerifyVoteExtensionResponse, error) {
	return &abcitypes.VerifyVoteExtensionResponse{}, nil
}
