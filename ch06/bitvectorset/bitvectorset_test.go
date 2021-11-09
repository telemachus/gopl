package bitvectorset_test

import (
	"bitvectorset"
	"fmt"
	"testing"
)

func sameInts(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, item := range a {
		if item != b[i] {
			return false
		}
	}
	return true
}

func TestElems(t *testing.T) {
	a := bitvectorset.IntSet{}

	t.Run("empty set => empty slice of ints", func(t *testing.T) {
		expected := []int{}
		actual := a.Elems()

		if !sameInts(expected, actual) {
			t.Errorf("expected %v; actual %v", expected, actual)
		}
	})

	t.Run("single item set => single item slice of ints", func(t *testing.T) {
		expected := []int{1}
		a.Add(1)
		actual := a.Elems()

		if !sameInts(expected, actual) {
			t.Errorf("expected %v; actual %v", expected, actual)
		}
	})

	t.Run("multi item set => multi item slice of ints", func(t *testing.T) {
		expected := []int{1, 2, 3, 4}
		a.AddAll(1, 2, 3, 4)
		actual := a.Elems()

		if !sameInts(expected, actual) {
			t.Errorf("expected %v; actual %v", expected, actual)
		}
	})
}

func TestString(t *testing.T) {
	t.Run("empty set", func(t *testing.T) {
		a := bitvectorset.IntSet{}
		expected := "{}"
		actual := fmt.Sprintf("%s", &a)
		if expected != actual {
			t.Errorf("expected %q; actual %q", expected, actual)
		}
	})

	t.Run("one item set", func(t *testing.T) {
		a := bitvectorset.IntSet{}
		a.Add(6)
		expected := "{6}"
		actual := fmt.Sprintf("%s", &a)
		if expected != actual {
			t.Errorf("expected %q; actual %q", expected, actual)
		}
	})

	t.Run("three item set", func(t *testing.T) {
		a := bitvectorset.IntSet{}
		a.Add(6)
		a.Add(2)
		a.Add(7)
		expected := "{2, 6, 7}"
		actual := fmt.Sprintf("%s", &a)
		if expected != actual {
			t.Errorf("expected %q; actual %q", expected, actual)
		}
	})
}

func TestLen(t *testing.T) {
	a := bitvectorset.IntSet{}

	t.Run("Len reports 0 for empty set", func(t *testing.T) {
		if a.Len() != 0 {
			t.Errorf("Len should say 0 for this set: %s", &a)
		}
	})

	a.Add(1)
	t.Run("Len reports 1 for a one-item set", func(t *testing.T) {
		if a.Len() != 1 {
			t.Errorf("Len should say 1 for this set: %s", &a)
		}
	})

	a.AddAll(2, 3, 4)
	t.Run("Len reports the correct length for sets", func(t *testing.T) {
		if a.Len() != 4 {
			t.Errorf("Len should say 4 for this set: %s", &a)
		}
	})
}

func TestAdd(t *testing.T) {
	a := bitvectorset.IntSet{}
	a.Add(3)
	a.Add(2)

	b := bitvectorset.IntSet{}
	b.Add(2)
	b.Add(3)
	b.Add(3)

	if !b.Same(&a) {
		t.Errorf("expected these two sets to be the same: %s, %s", &b, &a)
	}
}

func TestAddAll(t *testing.T) {
	a, b := bitvectorset.IntSet{}, bitvectorset.IntSet{}
	a.Add(1)
	a.Add(2)
	a.Add(3)
	b.AddAll(1, 2, 3)

	if !b.Same(&a) {
		t.Errorf("expected these two sets to be the same: %s, %s", &b, &a)
	}
}

func TestHas(t *testing.T) {
	a, b := bitvectorset.IntSet{}, bitvectorset.IntSet{}
	b.AddAll(1, 52, 1024)

	t.Run("Has on empty set returns false", func(t *testing.T) {
		if a.Has(0) {
			t.Errorf("the empty set does not have anything: %s", &a)
		}
	})

	t.Run("Has finds items in a set", func(t *testing.T) {
		for _, item := range []int{1, 52, 1024} {
			if !b.Has(item) {
				t.Errorf("set %s should have %d", &b, item)
			}
		}
	})
}

func TestRemove(t *testing.T) {
	a := bitvectorset.IntSet{}
	a.AddAll(1, 2, 526)

	t.Run("Remove removes item", func(t *testing.T) {
		a.Remove(1)
		if a.Has(1) {
			t.Errorf("this set should not have 1: %s", &a)
		}
	})

	t.Run("Remove leaves other items", func(t *testing.T) {
		if !a.Has(2) || !a.Has(526) {
			t.Errorf("this set should still have 2 and 526: %s", &a)
		}
	})

	a.Remove(526)
	t.Run("Remove doesn’t mess with earlier items", func(t *testing.T) {
		if a.Has(526) || !a.Has(2) {
			t.Errorf("this set should have only 2: %s", &a)
		}
	})
}

func TestClear(t *testing.T) {
	a, b, c := bitvectorset.IntSet{}, bitvectorset.IntSet{}, bitvectorset.IntSet{}
	b.Add(1)
	c.AddAll(1, 2, 3, 4, 5, 6, 7, 8, 9)

	t.Run("Clear should be harmless on an empty set", func(t *testing.T) {
		a.Clear()
		if a.Len() != 0 {
			t.Errorf("this set should be empty: %s", &a)
		}
	})

	t.Run("Clear should leave a set empty", func(t *testing.T) {
		b.Clear()
		if b.Len() > 0 {
			t.Errorf("this set should be empty: %s", &b)
		}
	})

	t.Run("Clear should leave empty sets equivalent", func(t *testing.T) {
		a.Clear()
		c.Clear()
		if !a.Same(&c) {
			t.Errorf("these sets should be identical: %s, %s", &a, &c)
		}
	})
}

func TestCopy(t *testing.T) {
	a := bitvectorset.IntSet{}

	t.Run("Copy works on empty sets", func(t *testing.T) {
		b := a.Copy()
		if !a.Same(b) {
			t.Errorf("these should both be empty sets: %s, %s", &a, b)
		}
	})

	t.Run("Copy copies a one-item set", func(t *testing.T) {
		a.Add(1)
		b := a.Copy()
		if b.Len() != 1 || !b.Has(1) {
			t.Errorf("this should be a one-item set containing 1: %s", b)
		}
	})

	t.Run("Copy copies a three-item set", func(t *testing.T) {
		a.AddAll(3, 2, 1)
		b := a.Copy()
		if b.Len() != 3 || !b.Has(1) || !b.Has(2) || !b.Has(3) {
			t.Errorf("this should be the set {1, 2, 3}: %s", b)
		}
	})
}

func TestUnionWith(t *testing.T) {
	a, b := bitvectorset.IntSet{}, bitvectorset.IntSet{}
	a.AddAll(1, 2, 3, 567, 987, 1023)
	b.AddAll(4, 5, 6, 678, 889, 1101)
	a.UnionWith(&b)

	t.Run("Union combines sets", func(t *testing.T) {
		for _, item := range []int{1, 2, 3, 4, 5, 6, 567, 678, 889, 987, 1023, 1101} {
			if !a.Has(item) {
				t.Errorf("%s should contain %d", &a, item)
			}
		}
	})

	t.Run("Union leaves argument set as is", func(t *testing.T) {
		for _, item := range []int{4, 5, 6, 678, 889, 1101} {
			if !b.Has(item) {
				t.Errorf("%s should contain %d", &b, item)
			}
		}

		for _, item := range []int{1, 2, 3, 567, 987, 1023} {
			if b.Has(item) {
				t.Errorf("%s should not contain %d", &b, item)
			}
		}
	})
}

func TestSame(t *testing.T) {
	a := bitvectorset.IntSet{}
	b := bitvectorset.IntSet{}
	c := bitvectorset.IntSet{}
	a.Add(1)
	b.Add(1)
	c.Add(2)

	t.Run("Same should report identical sets", func(t *testing.T) {
		if !a.Same(&b) {
			t.Errorf("%s.Same(&%s) should be true", &a, &b)
		}
	})

	t.Run("Same should report non-identical sets", func(t *testing.T) {
		if a.Same(&c) {
			t.Errorf("%s.Same(&%s) should be false", &a, &c)
		}
	})
}

func TestIntersectWith(t *testing.T) {

	t.Run("IntersectWith clears items not in both sets", func(t *testing.T) {
		b := bitvectorset.IntSet{}
		c := bitvectorset.IntSet{}
		b.AddAll(1, 2, 3, 4, 5)
		c.AddAll(6, 7, 8, 9, 10)

		b.IntersectWith(&c)
		if b.Len() > 0 {
			t.Errorf("%s should now be empty", &b)
		}
	})

	t.Run("IntersectWith keeps only items in both sets", func(t *testing.T) {
		a := bitvectorset.IntSet{}
		b := bitvectorset.IntSet{}
		a.AddAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		b.AddAll(1, 2, 3, 4, 5)

		a.IntersectWith(&b)
		if !a.Same(&b) {
			t.Errorf("%s and %s should now be the same", &a, &b)
		}
	})

	t.Run("IntersectWith keeps shared items regardless of order", func(t *testing.T) {
		a := bitvectorset.IntSet{}
		c := bitvectorset.IntSet{}
		a.AddAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		c.AddAll(6, 7, 8, 9, 10)

		a.IntersectWith(&c)
		if !a.Same(&c) {
			t.Errorf("%s and %s should now be the same", &a, &c)
		}
	})
}

func TestDifferenceWith(t *testing.T) {
	t.Run("a.DifferenceWith(b) a>b, removes items in b from a", func(t *testing.T) {
		a := bitvectorset.IntSet{}
		b := bitvectorset.IntSet{}
		a.AddAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		b.AddAll(3, 4, 5)

		a.DifferenceWith(&b)
		if a.Len() != 7 {
			t.Errorf("%s should now have seven items", &a)
		}

		for _, item := range []int{1, 2, 6, 7, 8, 9, 10} {
			if !a.Has(item) {
				t.Errorf("%s should have %d", &a, item)
			}
		}
	})

	t.Run("a.DifferenceWith(b) b>a, removes items in b from a", func(t *testing.T) {
		a := bitvectorset.IntSet{}
		b := bitvectorset.IntSet{}
		a.AddAll(1, 2, 3, 4)
		b.AddAll(3, 4, 5, 6, 7, 8, 9)

		a.DifferenceWith(&b)
		if a.Len() != 2 {
			t.Errorf("%s should now have two items", &a)
		}

		for _, item := range []int{1, 2} {
			if !a.Has(item) {
				t.Errorf("%s should have %d", &a, item)
			}
		}
	})
}

func TestSymmetricDifferenceWith(t *testing.T) {
	t.Run("a.SymmetricDifferenceWith(b) yields symmetric difference of a and b", func(t *testing.T) {
		a := bitvectorset.IntSet{}
		b := bitvectorset.IntSet{}
		a.AddAll(4, 5, 6, 7, 8, 9, 10)
		b.AddAll(1, 2, 3, 4, 5, 6)

		a.SymmetricDifferenceWith(&b)
		if a.Len() != 7 {
			t.Errorf("%s should now have seven items", &a)
		}

		for _, item := range []int{1, 2, 3, 7, 8, 9, 10} {
			if !a.Has(item) {
				t.Errorf("%s should have %d", &a, item)
			}
		}
	})

	t.Run("b.SymmetricDifferenceWith(a) yields symmetric difference of a and b", func(t *testing.T) {
		a := bitvectorset.IntSet{}
		b := bitvectorset.IntSet{}
		a.AddAll(4, 5, 6, 7, 8, 9, 10, 11, 12)
		b.AddAll(1, 2, 3, 4, 5, 6)

		b.SymmetricDifferenceWith(&a)
		if b.Len() != 9 {
			t.Errorf("%s should now have eight items", &a)
		}

		for _, item := range []int{1, 2, 3, 7, 8, 9, 10, 11, 12} {
			if !b.Has(item) {
				t.Errorf("%s should have %d", &b, item)
			}
		}
	})

	t.Run("SymmetricDifferenceWith on identical sets yields an empty set", func(t *testing.T) {
		a := bitvectorset.IntSet{}
		b := bitvectorset.IntSet{}
		a.AddAll(4, 5, 6)
		b.AddAll(4, 5, 6)

		b.SymmetricDifferenceWith(&a)
		if b.Len() != 0 {
			t.Errorf("%s should now be empty", &b)
		}
	})

	t.Run("SymmetricDifferenceWith a>b, one difference earlier", func(t *testing.T) {
		a := bitvectorset.IntSet{}
		b := bitvectorset.IntSet{}
		a.AddAll(1, 4, 5, 6)
		b.AddAll(4, 5, 6)

		b.SymmetricDifferenceWith(&a)
		if b.Len() != 1 || !b.Has(1) {
			t.Errorf("%s should now contain only 1", &b)
		}
	})

	t.Run("SymmetricDifferenceWith a>b, one difference later", func(t *testing.T) {
		a := bitvectorset.IntSet{}
		b := bitvectorset.IntSet{}
		a.AddAll(4, 5, 6, 7)
		b.AddAll(4, 5, 6)

		b.SymmetricDifferenceWith(&a)
		if b.Len() != 1 || !b.Has(7) {
			t.Errorf("%s should now contain only 7", &b)
		}
	})

	t.Run("SymmetricDifferenceWith b>a, one difference earlier", func(t *testing.T) {
		a := bitvectorset.IntSet{}
		b := bitvectorset.IntSet{}
		a.AddAll(4, 5, 6)
		b.AddAll(1, 4, 5, 6)

		b.SymmetricDifferenceWith(&a)
		if b.Len() != 1 || !b.Has(1) {
			t.Errorf("%s should now contain only 1", &b)
		}
	})

	t.Run("SymmetricDifferenceWith b>a, one difference later", func(t *testing.T) {
		a := bitvectorset.IntSet{}
		b := bitvectorset.IntSet{}
		a.AddAll(4, 5, 6)
		b.AddAll(4, 5, 6, 7)

		b.SymmetricDifferenceWith(&a)
		if b.Len() != 1 || !b.Has(7) {
			t.Errorf("%s should now contain only 7", &b)
		}
	})
}
