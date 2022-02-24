// Code generated by "json-ice --type=PeerDetails"; DO NOT EDIT.

package result

func MarshalPeerDetailsAsJSON(s *PeerDetails) ([]byte, error) {
	buff := make([]byte, 1, 54)
	buff[0] = '{'
	buff = append(buff, "\"iceUsernameFragment\":"...)
	if marshaled, err := s.IceUsernameFragment.MarshalJSON(); err != nil {
		return nil, err
	} else {
		buff = append(buff, marshaled...)
	}

	buff = append(buff, ',')
	buff = append(buff, "\"icePassword\":"...)
	if marshaled, err := s.IcePassword.MarshalJSON(); err != nil {
		return nil, err
	} else {
		buff = append(buff, marshaled...)
	}

	buff = append(buff, ',')
	if buff[len(buff)-1] == ',' {
		buff[len(buff)-1] = '}'
	} else {
		buff = append(buff, '}')
	}
	return buff, nil
}
