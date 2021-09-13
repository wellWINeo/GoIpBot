package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"

	"github.com/wellWINeo/GoIpBot"
)

const baseUrl = "http://api.ipstack.com"

type IpStackRepository struct {
	token string
}

func NewIpStack(token string) *IpStackRepository {
	return &IpStackRepository{token: token}
}

func (i *IpStackRepository) GetInfo(ip net.IP) (GoIpBot.IpInfo, error) {
	requestURL := fmt.Sprintf("%s/%s?access_key=%s", baseUrl, ip, i.token)
	resp, err := http.Get(requestURL)
	if err != nil {
		return GoIpBot.IpInfo{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return GoIpBot.IpInfo{}, errors.New("Bad response from API")
	}

	defer resp.Body.Close()

	var info GoIpBot.IpInfo
	err = json.NewDecoder(resp.Body).Decode(&info)

	return info, err
}
