package forkchoice

import (
	"testing"

	"github.com/theQRL/qrysm/config/params"
	"github.com/theQRL/qrysm/consensus-types/blocks"
	"github.com/theQRL/qrysm/testing/require"
	"github.com/theQRL/qrysm/testing/util"
)

func TestBuilderTick(t *testing.T) {
	st, err := util.NewBeaconStateZond()
	require.NoError(t, err)
	blk, err := blocks.NewSignedBeaconBlock(util.NewBeaconBlockZond())
	require.NoError(t, err)
	builder := NewBuilder(t, st, blk)
	builder.Tick(t, 10)

	require.Equal(t, int64(10), builder.lastTick)
}

func TestBuilderInvalidBlock(t *testing.T) {
	st, err := util.NewBeaconStateZond()
	require.NoError(t, err)
	genesisBlk, err := blocks.NewSignedBeaconBlock(util.NewBeaconBlockZond())
	require.NoError(t, err)
	parentRoot, err := genesisBlk.Block().HashTreeRoot()
	require.NoError(t, err)
	blkPb := util.NewBeaconBlockZond()
	blkPb.Block.Slot = 1
	blkPb.Block.ParentRoot = parentRoot[:]
	blk, err := blocks.NewSignedBeaconBlock(blkPb)
	require.NoError(t, err)
	builder := NewBuilder(t, st, genesisBlk)
	builder.Tick(t, int64(params.BeaconConfig().SecondsPerSlot))
	status := "INVALID"
	require.NoError(t, builder.SetPayloadStatus(&MockEngineResp{Status: &status}))
	builder.InvalidBlock(t, blk)
}

/*
func TestPoWBlock(t *testing.T) {
	st, err := util.NewBeaconStateZond()
	require.NoError(t, err)
	blk, err := blocks.NewSignedBeaconBlock(util.NewBeaconBlockZond())
	require.NoError(t, err)
	builder := NewBuilder(t, st, blk)
	builder.PoWBlock(&qrlpb.PowBlock{BlockHash: []byte{1, 2, 3}})

	require.Equal(t, 1, len(builder.execMock.powBlocks))
}
*/
