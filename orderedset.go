package orderedSet





type orderedSet struct{
	Set []interface{}
	size int
	cmp func(a, b interface{}) int
}

 


func NewSet(compare func(a, b interface{}) int, elems ...interface{}) orderedSet {
	


	st := orderedSet{}
	st.cmp = compare
	st.Set = make([]interface{}, 0)
	st.size = 0
	
	for _, v := range elems {
		st.Insert(v)
	}

	return st
}
func (st *orderedSet)Size()int{
	return st.size
}

func (st *orderedSet)isEmpty() bool{
	return (st.size == 0)
}

func (st *orderedSet) Lower_bound(elem interface{}) int {
	l := 0
	r := st.size-1

	for {
		if r <= l {
			break
		}
		mid := l + (r-l)/2
		if mid >= st.size{
			break
		}
		if st.cmp(elem, (st.Set)[mid]) != 1 {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if	l <st.size&&st.cmp((st.Set)[l],elem)==-1 {
		l++
	 }
	return l
}

func (st orderedSet) Bsearch(elem interface{}) int {
	l := 0
	r := st.size - 1
	for {
		mid := (r + l) / 2
		if st.cmp(st.Set[mid], elem) == 0 {
			return mid
		}
		if l>=r{
			break
		}

		if st.cmp(elem,st.Set[mid]) == -1 {
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

	if i >= st.size{
		(st.Set) = append(st.Set, elem)
		st.size++
		return true
	}
	
	if i == 0 && (st.cmp(elem, (st.Set)[0]) != 0) {
		
		st.Set = append([]interface{}{elem}, (st.Set)[0:]...)
		st.size++
		return true
	}
		
	if st.cmp(elem, (st.Set)[i]) != 0 {
		(st.Set)= append((st.Set)[0:i+1], (st.Set)[i:]...)
		(st.Set)[i] = elem
		st.size++
	return true
}
	return false
}
func (st *orderedSet) Remove(elem interface{}) bool {
	i := (*st).Bsearch(elem)
	if i == -1 {
		return false
	} else {
		(st.Set) = append((st.Set)[:i], (st.Set)[i+1:]...)
		st.size--
		return true
	}
	

}

//helper functions
