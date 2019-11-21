package main

type errorBody struct {
	Err string `json:"error, omitempty"`
}

//func CustomHTTPError(ctx context.Context, _ *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {
//	const fallback = `{"error": "failed to marshal error message"}`
//
//	w.Header().Set("Content-type", marshaler.ContentType())
//	w.WriteHeader(runtime.HTTPStatusFromCode(status.Code(err)))
//	jErr := json.NewEncoder(w).Encode(errorBody{
//		Err: err.Error(),
//	})
//
//	if jErr != nil {
//		w.Write([]byte(fallback))
//	}
//}