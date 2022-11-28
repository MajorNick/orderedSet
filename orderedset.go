package orderedSet





type orderedSet []interface{}

var cmp func(a, b interface{}) int


func NewSet(compare func(a, b interface{}) int, elems ...interface{}) orderedSet {
	cmp = compare

	st := make(orderedSet, 0)
	for _, v := range elems {
		st.Insert(v)
	}

	return st
}
func (st *orderedSet)Size()int{
	return len(*st)
}

func (st *orderedSet)isEmpty() bool{
	return (len(*st) == 0)
}

func (st *orderedSet) Lower_bound(elem interface{}) int {
	l := 0
	r := len(*st)-1

	for {
		if r <= l {
			break
		}
		mid := l + (r-l)/2
		if mid >= len(*st){
			break
		}
		if cmp(elem, (*st)[mid]) != 1 {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if	l <len(*st)&&cmp((*st)[l],elem)==-1 {
		l++
	 }
	return l
}

func (st orderedSet) Bsearch(elem interface{}) int {
	l := 0
	r := len(st) - 1
	for {
		mid := (r + l) / 2
		if cmp(st[mid], elem) == 0 {
			return mid
		}
		if l>=r{
			break
		}

		if cmp(elem,st[mid]) == -1 {
			r = mid - 1
			
		}else{
			l = mid + 1
		}
	}

	return -1
}
func (st orderedSet) Contains(elem interface{}) bool {
	return st.Bsearch(elem) != -1
}
func (st *orderedSet) Insert(elem interface{}) bool {

	
	i := st.Lower_bound(elem)

	if i >= len(*st) {
		*st = append(*st, elem)

		return true
	}
	
	if i == 0 && (cmp(elem, (*st)[0]) != 0) {
		
		*st = append([]interface{}{elem}, (*st)[0:]...)
		
		return true
	}
		
	if cmp(elem, (*st)[i]) != 0 {
		(*st)= append((*st)[0:i+1], (*st)[i:]...)
		(*st)[i] = elem

	return true
}
	return false
}
func (st *orderedSet) Remove(elem interface{}) bool {
	i := (*st).Bsearch(elem)
	if i == -1 {
		return false
	} else {
		(*st) = append((*st)[:i], (*st)[i+1:]...)

		return true
	}
	

}

//helper functions
