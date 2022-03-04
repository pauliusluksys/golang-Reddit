package servicesMail

//
//type VerificationDataKey struct{}
//type VerificationData struct {
//	Email     string               `json:"email" validate:"required" sql:"email"`
//	Code      string               `json:"code" validate:"required" sql:"code"`
//	ExpiresAt time.Time            `json:"expiresat" sql:"expiresat"`
//	Type      VerificationDataType `json:"type" sql:"type"`
//}
//
//func VerifyMail(w http.ResponseWriter, r *http.Request) {
//
//	w.Header().Set("Content-Type", "application/json")
//	verificationData := r.Context().Value(VerificationDataKey{}).(VerificationData)
//
//}
