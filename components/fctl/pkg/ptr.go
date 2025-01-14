package fctl

func BoolToString(v bool) string {
	if !v {
		return "No"
	}
	return "Yes"
}

func BoolPointerToString(v *bool) string {
	if v == nil {
		return "No"
	}
	return BoolToString(*v)
}

func StringPointerToString(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}
