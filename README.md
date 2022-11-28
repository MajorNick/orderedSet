# Ordered Set

## Ordered Set from Classic STL.

### Writen using empty interfaces. So it can called Generics without actual Generics.

### Pros: every "Big" method works in O(log(n)) complexity and can be used on all types

### Cons: Regardless of whether it is primitive or not, you have to give NewSet()  Compare function.
### And Also after inserting struct type in set and you need type assertion after getting element  from Set.

### Example compare function for Integers:

 ```
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
```
