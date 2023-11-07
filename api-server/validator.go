package api_server

//func use {
//	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
//		vs := map[string]validator.Func{
//			"alphanumunicode": alphanumunicode,
//			"nolinebreak":     nolinebreak,
//		}
//
//		for k, fn := range vs {
//			if err := v.RegisterValidation(k, fn); err != nil {
//				return err
//			}
//		}
//	}
//}
//
//func alphanumunicode(fl validator.FieldLevel) bool {
//	allowed := map[rune]bool{
//		'·': true,
//		'/': true,
//		'-': true,
//		'_': true,
//	}
//
//	field := fl.Field().String()
//	for _, r := range field {
//		if !unicode.IsLetter(r) && !unicode.IsNumber(r) && !unicode.IsSpace(r) && !allowed[r] {
//			return false
//		}
//	}
//	return true
//}
//
//func nolinebreak(fl validator.FieldLevel) bool {
//	field := fl.Field().String()
//	return !strings.Contains(field, "\n")
//}
