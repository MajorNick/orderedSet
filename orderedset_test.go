package orderedSet

import (

	"testing"
)

func TestSize(t *testing.T){
	cmp := func (a,b interface{}) int {
		a1 := a.(int)
		b1 := b.(int)
		if a1 == b1 {
			return 0
		}
		if(a1>b1){
			return 1
		}
		return -1
	}
	st := NewSet(cmp)
	for i:=1;i<27;i++{
		st.Insert(i)
		if st.Size() != i{
			t.Fatalf("Expected Size %d but Found %d!",i,st.Size())
		}
	}



	for i:=1;i<27;i++{
		st.Remove(i)
		if st.Size() != 26-i{
			t.Fatalf("Expected Size %d but Found %d!",26-i,st.Size())
		}
	}

}
func TestOnStrings(t *testing.T){
	cmp := func (a,b interface{}) int {
		a1 := a.(string)
		b1 := b.(string)
		if a1 == b1 {
			return 0
		}
		if(a1>b1){
			return 1
		}
		return -1
	}
	st := NewSet(cmp)
	tmp := []string{}
	for i:='a';i<='z';i++{
		st.Insert(string(i))
		tmp = append(tmp,string(i))
		
	}
	for i:=0;i<26;i++{
		if tmp[i] != st[i]{
			t.Fatalf("Set Isn't Sorted! expected %s, found %s!\n",tmp[i],st[i])
		}
	}


	for i:='a';i<='z';i++{
		
		k:=st.Bsearch(string(i))
		
		if k == -1||string(i) != st[k]{
			t.Fatalf("Wrong Answer in Bsearch! expected %s, found %s!\n",tmp[i],st[i])
		}
	}
}
func TestOnStructs(t *testing.T){
	type tp struct{
		name string
		surname string
		id int
	}
	cmp := func (a,b interface{}) int {
		a1 := a.(tp)
		b1 := b.(tp)
		if a1.id == b1.id {
			return 0
		}
		if(a1.id>b1.id){
			return 1
		}
		return -1
	}
	
	st := NewSet(cmp)
	for i:=9;i>=1;i--{
		tmp := tp{"","",i}
		st.Insert(tmp)
	}

	for i:=1;i<10;i++{
		
		if st[i-1].(tp).id!= i{
			t.Fatalf("Set Isn't Sorted! expected %v, found %v!\n",tp{"","",i},st[i])
		}
	}

	for i:=1;i<10;i++{
		
		k:=st.Bsearch(tp{"","",i})
		tmp := tp{"","",i}
		if tmp != st[k]{
			t.Fatalf("Wrong Answer in Bsearch! expected %v, found %v!\n",tp{"","",i},st[i])
		}
	}

}
func TestLower_Bound(t *testing.T) {
	cmp := func (a,b interface{}) int {
		a1 := a.(int)
		b1 := b.(int)
		if a1 == b1 {
			return 0
		}
		if(a1>b1){
			return 1
		}
		return -1
	}
	
	st := NewSet(cmp)
	slc := []int{1,5,78,8,6,4,3,7}
	sorted := []int{1,3,4,5,6,7,8,78}
	for i:=0;i<len(sorted);i++{
		st.Insert(slc[i])
	}
	
	for i:=0;i<len(sorted);i++{
		
		if sorted[i]!=st[i]{
			t.Fatalf("Set Isn't Sorted! expected %d, found %d!\n",sorted[i],st[i])
		}
	} 

	k := st.Lower_bound(9)
	if k != 7{
		t.Fatalf("Wrong Result From Lower_Bound! expected 7, found %d",k)
	}
	k = st.Lower_bound(2)
	if k != 1{
		t.Fatalf("Wrong Result From Lower_Bound! expected 1, found %d",k)
	}
	k = st.Lower_bound(0)
	if k != 0{
		t.Fatalf("Wrong Result From Lower_Bound! expected 0, found %d",k)
	}

}
