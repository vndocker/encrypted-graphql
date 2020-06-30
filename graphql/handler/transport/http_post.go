package transport

import (
	"fmt"
	"mime"
	"net/http"

	"github.com/vndocker/encrypted-graphql/graphql"
)

// POST implements the POST side of the default HTTP transport
// defined in https://github.com/APIs-guru/graphql-over-http#post
type POST struct{}

var _ graphql.Transport = POST{}

func (h POST) Supports(r *http.Request) bool {
	if r.Header.Get("Upgrade") != "" {
		return false
	}

	mediaType, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return false
	}

	supportedJson := r.Method == "POST" && mediaType == "application/json"

	encryptWith := r.Header.Get("X-Encrypted-With")

	if encryptWith == "RSA" && supportedJson {
		return true
	}
	if encryptWith != "" && supportedJson {
		return false
	}

	return supportedJson
}


func (h POST) Do(w http.ResponseWriter, r *http.Request, exec graphql.GraphExecutor) {
	w.Header().Set("Content-Type", "application/json")

	var params *graphql.RawParams
	start := graphql.Now()

	// fmt.Printf("Content Input Transport: %v", r.Body)

	//get private key
	//priv, err := pki.GetPrivateKey()
	//if err != nil {
	//	log.Fatalf("Error: %s\n", err)
	//}
	//
	//decrypted := pki.Decrypt(requestBody.CONTENT, priv)
	//fmt.Printf("decrypted: %s:: \n", decrypted)

	if err := jsonDecode(r.Body, &params); err != nil { // r.Body = ShouldDecryptedString
		w.WriteHeader(http.StatusBadRequest)
		writeJsonErrorf(w, "json body could not be decoded: "+err.Error())
		return
	}


	params.ReadTime = graphql.TraceTiming{
		Start: start,
		End:   graphql.Now(),
	}

	rc, err := exec.CreateOperationContext(r.Context(), params)

	// fmt.Printf("Prepare output: Params: %v, RC: %v", params, rc)

	if err != nil {
		w.WriteHeader(statusFor(err))
		resp := exec.DispatchError(graphql.WithOperationContext(r.Context(), rc), err)
		writeJson(w, resp)
		return
	}
	ctx := graphql.WithOperationContext(r.Context(), rc)
	responses, ctx := exec.DispatchOperation(ctx, rc)

	fmt.Printf("Output: responses: %v", responses)

	writeJson(w, responses(ctx))
}
