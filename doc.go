// Copyright 2015 me. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Dynamic Connectivity and its Union Find implementation are
inspired by Algorithms Lecture Slide by ROBERT SEDGEWICK &
KEVIN WAYNE.

The package implements compatible Union-Find API as illustrated
in the lecture slides. The Goal is to have

  a efficient data structure for large number of objects

  and large number of operations

A dynamic connectivity client (test client) consumes the API
and construct a connectivity object
*/
package connectivity

/*
History of implementation

First implementation decision for a FindQuery() or 'Connected()' is to have
an array combed. But that's obviously O(n) complexity if the
pocket grows larger and larger, not to mention when the container
of the pockets holds huge number of large pockets the Query will
take a long time.

NOTE(nilbot): But the first intuition was not array but a Set, it was however
rejected because I was comfortable only with array and slice in goalng.
Implementing a set will need a map, and I thought map is quite heavy weight,
oblivious to the fact that map[int]struct{} costs almost nothing!

Following that implementation we got a FindQuery of O(n^2) and
an Union of O(n^2). So bad.

*/
