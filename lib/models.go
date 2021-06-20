package lib

type StakeEntryStats struct {
}

const HashSizeBytes = 32

type BlockHash [HashSizeBytes]byte
type StakeEntry struct {
}

type PostEntry struct {
	PostHash                 *BlockHash
	PosterPublicKey          []byte
	ParentStakeID            []byte
	Body                     []byte
	RecloutedPostHash        *BlockHash
	IsQuotedReclout          bool
	CreatorBasisPoints       uint64
	StakeMultipleBasisPoints uint64
	ConfirmationBlockHeight  uint32
	TimestampNanos           uint64
	IsHidden                 bool
	StakeEntry               *StakeEntry
	LikeCount                uint64
	RecloutCount             uint64
	QuoteRecloutCount        uint64
	DiamondCount             uint64
	stakeStats               *StakeEntryStats
	isDeleted                bool
	CommentCount             uint64
	IsPinned                 bool
	PostExtraData            map[string][]byte
}
