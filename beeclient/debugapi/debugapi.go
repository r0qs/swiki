// Copyright 2020 Ethersphere. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// This code is based on the beekeeper beeclient api

package debugapi

import (
	"net/url"
	"swiki/httpclient"
)

type DebugAPI struct {
	*httpclient.Client
	Node    *NodeService
	Postage *PostageService
}

func NewDebugAPI(beeURL *url.URL, o *httpclient.ClientOptions) (*DebugAPI, error) {
	httpc, err := httpclient.NewClient(beeURL, o)
	if err != nil {
		return nil, err
	}
	c := &DebugAPI{httpc, nil, nil}
	c.Node = newNodeService(c)
	c.Postage = newPostageService(c)
	return c, nil
}