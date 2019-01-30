// Copyright (c) 2015-present Corey Scott (www.sage42.com), All Rights Reserved.
//
// NOTICE: All information contained herein is, and remains the property of Corey Scott.
// The intellectual and technical concepts contained herein are confidential, proprietary and controlled by Corey Scott
// and may be covered by patents, patents in process, and are protected by trade secret or copyright law.
//
// You are strictly forbidden to copy, download, store (in any medium), transmit, disseminate, adapt or change this
// material in any way unless prior written permission is obtained from Corey Scott.
// Access to the source code contained herein is hereby forbidden to anyone except explicit written consent and subject
// to binding Confidentiality and Non-disclosure agreements explicitly covering such access.
//
// The copyright notice above does not evidence any actual or intended publication or disclosure of this source code,
// which includes information that is confidential and/or proprietary, and is a trade secret, of Corey Scott.
//
// ANY REPRODUCTION, MODIFICATION, DISTRIBUTION, PUBLIC PERFORMANCE, OR PUBLIC DISPLAY OF OR THROUGH USE OF THIS SOURCE
// CODE WITHOUT THE EXPRESS WRITTEN CONSENT OF COREY SCOTT IS STRICTLY PROHIBITED, AND IN VIOLATION OF APPLICABLE LAWS
// AND INTERNATIONAL TREATIES. THE RECEIPT OR POSSESSION OF THIS SOURCE CODE AND/OR RELATED INFORMATION DOES NOT CONVEY
// OR IMPLY ANY RIGHTS TO REPRODUCE, DISCLOSE OR DISTRIBUTE ITS CONTENTS, OR TO MANUFACTURE, USE, OR SELL ANYTHING
// THAT IT MAY DESCRIBE, IN WHOLE OR IN PART.

package v1

import (
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"net/http"
)

func GetUserEndpoint(resp http.ResponseWriter, req *http.Request) {
	// get and check inputs
	ID, err := getRequestedID(req)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	// load requested data
	user, err := loadUser(ID)
	if err != nil {
		// technical error
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	if user == nil {
		// user not found
		resp.WriteHeader(http.StatusNoContent)
		return
	}

	// prepare output
	switch req.Header.Get("Accept") {
	case "text/csv":
		outputAsCSV(resp, user)

	case "application/xml":
		outputAsXML(resp, user)

	case "application/json":
		fallthrough

	default:
		outputAsJSON(resp, user)
	}
}

func getRequestedID(req *http.Request) (int64, error) {
	// not yet implemented
	return 1, nil
}

func loadUser(ID int64) (*User, error) {
	// not yet implemented
	return &User{Name: "Fred", Phone: "0123456789"}, nil
}

func outputAsCSV(resp http.ResponseWriter, user *User) {
	writer := csv.NewWriter(resp)

	fields := []string{
		user.Name, user.Phone,
	}

	err := writer.Write(fields)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func outputAsXML(resp http.ResponseWriter, user *User) {
	encoder := xml.NewEncoder(resp)
	err := encoder.Encode(user)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func outputAsJSON(resp http.ResponseWriter, user *User) {
	encoder := json.NewEncoder(resp)
	err := encoder.Encode(user)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
}

type User struct {
	Name  string
	Phone string
}
