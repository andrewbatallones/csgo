package leetcode

func IsSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	// These will keep track of both of our trees
	// We don't need to have a "viewed" slice since we are not backtracking.
	// We can add the children here since we are checking the root node.
	// Why we need to add both even though they may be nil?
	// - It's an edge case that can happen if q.Left == p.Right and would produce a false-positive
	pQueue := [](*TreeNode){}
	qQueue := [](*TreeNode){}

	if p != nil {
		pQueue = append(pQueue, p.Left)
		pQueue = append(pQueue, p.Right)
	}

	if q != nil {
		qQueue = append(qQueue, q.Left)
		qQueue = append(qQueue, q.Right)
	}

	// Set currnet nodes
	pCurrent := p
	qCurrent := q

	// Essentially we can do dfs/bfs, but we'll check each node along the traversals
	// and if they don't match, return false
	for len(pQueue) > 0 && len(qQueue) > 0 {
		// If both are nil, we just pop from queue
		if pCurrent == nil && qCurrent == nil {
			pCurrent = pQueue[0]
			qCurrent = qQueue[0]

			pQueue = pQueue[1:]
			qQueue = qQueue[1:]
			continue
		}

		// Check values
		// We are going to break down the if statements for readablility, but can easily be combined
		// First check is if one or the other are nil (we checked prior for both nil)
		if pCurrent == nil || qCurrent == nil {
			return false
		}

		// Check values
		if pCurrent.Val != qCurrent.Val {
			return false
		}

		// Add children to queue
		pQueue = append(pQueue, pCurrent.Left)
		pQueue = append(pQueue, pCurrent.Right)
		qQueue = append(qQueue, qCurrent.Left)
		qQueue = append(qQueue, qCurrent.Right)

		// Assign new current
		pCurrent = pQueue[0]
		qCurrent = qQueue[0]

		// Pop off from slice
		pQueue = pQueue[1:]
		qQueue = qQueue[1:]
	}

	// After both queues are empty, if it didn't return false
	// we need to check if one of the lengths stopped early
	return len(pQueue) == len(qQueue)
}
