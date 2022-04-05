package response

import "gandi.icu/demo/model/system"

type SysReferralResponse struct {
	Referral system.SysReferral `json:"referral"`
}
