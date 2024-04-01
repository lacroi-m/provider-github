/*
Copyright 2019 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package clients

import (
	"net/http"

	github "github.com/google/go-github/github"
	// "github.com/lacroi-m-insta/provider-github/apis/projects/v1alpha1"
	// "github.com/lacroi-m-insta/provider-github/apis/v1beta1"
)

// Config provides github configurations for the Github client
type Config struct {
	Token              string
	BaseURL            string
	InsecureSkipVerify bool
	// AuthMethod         v1beta1.AuthType
}

// NewClient creates new Github Client with provided Github Configurations/Credentials.
func NewClient(c Config) *github.Client {
	var cl *github.Client
	var err error

	// TODO: configure client
	httpClient := &http.Client{}
	cl = github.NewClient(httpClient)

	/*
		if c.BaseURL != "" {
			options = append(options, github.WithBaseURL(c.BaseURL))
		}
		if c.InsecureSkipVerify {
			transport := cleanhttp.DefaultPooledTransport()
			if transport.TLSClientConfig == nil {
				transport.TLSClientConfig = &tls.Config{
					MinVersion: tls.VersionTLS12,
				}
			}
			transport.TLSClientConfig.InsecureSkipVerify = true
			httpclient := &http.Client{
				Transport: transport,
			}
			options = append(options, github.WithHTTPClient(httpclient))
		}

		switch c.AuthMethod {
		// TODO: implement BasicAuth structure with Username and Password
		//	case v1beta1.BasicAuth:
		// 		cl, err = github.NewBasicAuthClient(c.Username, c.Password, github.WithBaseURL(c.BaseURL))
		case v1beta1.JobToken:
			cl, err = github.NewJobClient(c.Token, options...)
		case v1beta1.OAuthToken:
			cl, err = github.NewOAuthClient(c.Token, options...)
		case v1beta1.PrivateToken:
			cl, err = github.NewClient(c.Token, options...)
		default:
			cl, err = github.NewClient(c.Token, options...)
		}
		if err != nil {
			panic(err)
		}
	*/
	if err != nil {
		panic(err)
	}

	return cl
}
