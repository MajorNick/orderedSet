package orderedSet

import(
	"fmt"
)



type OrderedSet struct{
	set []interface{}
	size int
	cmp func(a, b interface{}) int
}

 /*
	Getting Compare funmctions and slice of interface empty interface elements
	as Variadic parameters and adding Them to the set  
 */


func NewSet(compare func(a, b interface{}) int, elems ...interface{}) OrderedSet {
	


	st := OrderedSet{}
	st.cmp = compare
	st.set = make([]interface{}, 0)
	st.size = 0
	
	for _, v := range elems {
		st.Insert(v)
	}

	return st
}

func (st OrderedSet)Get(index int) interface{}{
	return st.set[index]
}

//Returns Current Number Of Elements in Set

func (st *OrderedSet)Size()int{
	return st.size
}
//Returns Current State of Set (usefull for "while" loop)
func (st *OrderedSet)isEmpty() bool{
	return (st.size == 0)
}

/* 
	Classic Lower_bound From STL, Algorithm is Similar To Binary search
*/
func (st *OrderedSet) Lower_bound(elem interface{}) int {
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
		if st.cmp(elem, st.Get(mid)) != 1 {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if	l <st.size&&st.cmp((st.set)[l],elem)==-1 {
		l++
	 }
	return l
}

/* 
	Classic Binary Search algorithm
*/

func (st OrderedSet) Bsearch(elem interface{}) int {
	l := 0
	r := st.size - 1
	for {
		mid := (r + l) / 2
		if st.cmp(st.Get(mid), elem) == 0 {
			return mid
		}
		if l>=r{
			break
		}

		if st.cmp(elem,st.Get(mid)) == -1 {
			r = mid - 1
			
		}else{
			l = mid + 1
		}
	}

	return -1
}

/*
	Contains function Using Binary search.
	Returns true if Element exists in Set
*/
func (st OrderedSet) Contains(elem interface{}) bool {
	return st.Bsearch(elem) != -1
}

/*
	Insert Function, Based on Lower_bound function.
	Lower Bound is giving me exact position of new Element
	in ordered Set.
*/

func (st *OrderedSet) Insert(elem interface{}) bool {

	
	i := st.Lower_bound(elem)

	if i >= st.size{
		(st.set) = append(st.set, elem)
		st.size++
		return true
	}
	
	if i == 0 && (st.cmp(elem, (st.set)[0]) != 0) {
		
		st.set = append([]interface{}{elem}, (st.set)[0:]...)
		st.size++
		return true
	}
		
	if st.cmp(elem, (st.set)[i]) != 0 {
		(st.set)= append((st.set)[0:i+1], (st.set)[i:]...)
		(st.set)[i] = elem
		st.size++
	return true
}
	return false
}

/*
	Finding element using Binary search,
	if element Exists, doing some slicing things 
	and and giving meat to the garbage collector
	 in the form of unnecessary  memory, which is removed 
	 from Set
*/
func (st *OrderedSet) Remove(elem interface{}) bool {
	i := (*st).Bsearch(elem)
	if i == -1 {
		return false
	} else {
		(st.set) = append((st.set)[:i], (st.set)[i+1:]...)
		st.size--
		return true
	}
	

}

// Classic ToString function

func (st OrderedSet) ToString()string{
	s := ""
	for i:=0;i<len(st.set);i++{
		s += fmt.Sprintf("%v ",st.set[i])
	}
	return s 
}

