package mfterror

var (
	ERR_UNKNOWN           = &MFTError{ErrorTypeInternal, "ERR_UNKNOWN", ""}
	ERR_CONNECTING_MF_API = &MFTError{ErrorTypeInternal, "ERR_CONNECTING_MF_API", ""}
	ERR_BAD_REQUEST       = &MFTError{ErrorTypeInvalidArgument, "ERR_BAD_REQUEST", ""}
)
