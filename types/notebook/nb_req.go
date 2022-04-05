package notebook

type AddAddressNoteBookReq struct {
	DeviceId     string `json:"device_id"`
	Name         string `json:"name"`
	AssetName    string `json:"asset_name"`
	Memo         string `json:"memo"`
	Address      string `json:"address"`
}

type UpdateAddressNoteBookReq struct {
	NbId int64 `json:"nb_id"`
	AddAddressNoteBookReq
}

type DelAddressNoteBookReq struct {
	NbId int64 `json:"nb_id"`
}


type QueryAddressNoteBookReq struct {
	DeviceId  string `json:"device_id"`
}

