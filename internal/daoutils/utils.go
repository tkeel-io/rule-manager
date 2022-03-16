package daoutils

/*


	这个package用于绕过service，对dao进行操作。实现一些特定的操作，并归档于这里。

*/

func inSlice(s []string, elem string) bool {
	for _, e := range s {
		if e == elem {
			return true
		}
	}
	return false
}
