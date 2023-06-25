# SharedBuffer

## Overview
The SharedBuffer project aims to provide a solution for managing shared buffers from contiguous memory in Go. It offers three different algorithms - counting, bitmap, and interval tree - to efficiently allocate and track memory buffers within a shared memory space.

By leveraging the SharedBuffer library in their Go applications, developers can efficiently manage shared buffers from contiguous memory using counting, bitmap, or interval tree-based algorithms, based on their specific requirements and performance considerations.

## Key Features:

Contiguous Memory Management: SharedBuffer facilitates the management of contiguous memory, allowing for efficient allocation and deallocation of memory buffers.

Three Algorithms: The project provides three algorithms to handle shared buffers:

* Counting: The counting algorithm utilizes a counter-based approach to track available memory buffers. It maintains a count of free buffers and efficiently determines available slots.

* Bitmap: The bitmap algorithm employs a bitmap data structure to represent the allocation status of memory buffers. It uses bitwise operations to manage and track free and allocated slots.

* Interval Tree: The interval tree algorithm organizes the shared memory space into a balanced tree structure. It efficiently handles buffer allocation and deallocation operations by storing intervals and searching for available slots within the tree.

## Installation

In order to use the library.
```
go get github.com/Giulio2002/sharedbuffer
```

## Usage

```go
buf, cancelFn := sharedbuffer.Make(4)
defer cancelFn()
```

By default the global shared buffer will have automatic locking and will be based on interval-trees rather than to the other algorithms.

### Non-global custom buffer

```go
s := NewConcurrentSharedBuffer(fsm.NewBitmapFreeSpaceManager(), management.NewMemoryBuffer())

buf, cancelFn := s.Make(4)
defer cancelFn()
```

You can use the following fsm implementation for the buffer:

* `fsm.NewBitmapFreeSpaceManager`: Bitmap based free space manager.
* `fsm.NewCountingFreeSpaceManager`: Counting based free space manager.
* `fsm.NewIntervalTreeFreeSpaceManager`: Interval balanced tree based free space manager.

### Avaiable Algorithms

### Counting Algorithm

In the counting algorithm, finding free slots involves scanning the contiguous memory space and checking the count of available buffers. The algorithm iterates over the memory slots until it encounters an available buffer. This linear search continues until a free slot is found.

#### Complexity


Finding Free Slots: O(n) - The counting algorithm requires a linear scan of the memory space to find a free slot. The time complexity is directly proportional to the number of slots in the memory.

### Bitmap Algorithm:

The bitmap algorithm efficiently finds free slots by performing bitwise operations on the bitmap. By examining the bits within the bitmap, the algorithm can identify available slots. It searches for the first unset bit (indicating a free slot) in the bitmap and returns the corresponding slot index.
**Note: Each slot is a word and each word has the same size so small data can cause internal fragmentation.*

#### Complexity

Finding Free Slots: O(N) - The bitmap algorithm requires scanning the bitmap, iterating over each bit, until an unset bit is encountered. The time complexity is directly proportional to the number of bits or slots represented by the bitmap


### Interval Tree Algorithm:
Finding free slots in the interval tree algorithm involves traversing the tree to locate an interval representing available memory buffers. The algorithm searches the tree for the smallest free interval that satisfies the required buffer size.

#### Complexity

Finding Free Slots: O(log n) - The interval tree algorithm has logarithmic time complexity for finding free slots. It performs a search in the interval tree, traversing nodes until it identifies a suitable free interval that can accommodate the requested buffer size.

