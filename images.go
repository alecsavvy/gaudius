package gaudius

import "fmt"

func (sdk *AudiusSdk) GetImage(cid string) ([]byte, error) {
	cn := sdk.Content
	dimensions := ImageDimensions("")
	res, err := cn.client.R().Get(fmt.Sprintf("/content/%s/%s.jpg", cid, dimensions))
	if err != nil {
		return nil, err
	}
	return res.Body(), nil
}
