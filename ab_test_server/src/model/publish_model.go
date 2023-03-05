package model

import (
	"bytes"
	"encoding/binary"
	common "line_china/common/model"
)

type PublishModel struct {
	AppName      string      `json:"app_name"`
	LayerName    string      `json:"layer_name"`
	ExpName      string      `json:"exp_name"`
	AppId        int         `json:"app_id"`
	DomainId     int         `json:"domain_id"`
	LayerId      int         `json:"layer_id"`
	ExpId        int         `json:"exp_id"`
	DomainWeight int         `json:"domain_weight"`
	ExpWeight    int         `json:"exp_weight"`
	ExpConfig    common.JSON `json:"exp_config"`
}

func (obj *PublishModel) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := binary.Write(buf, binary.BigEndian, obj); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
