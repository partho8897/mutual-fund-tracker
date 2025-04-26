package mfterror

type MFTErrorTypes uint32

const (
	ErrorTypeInternal         MFTErrorTypes = 0
	ErrorTypeInvalidArgument  MFTErrorTypes = 1
	ErrorTypeNotSupported     MFTErrorTypes = 2
	ErrorTypeInvalidOperation MFTErrorTypes = 3
)
