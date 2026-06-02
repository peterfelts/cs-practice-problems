package bruno

/*
	4. Lowest Common Ancestor of a Binary Tree
	Prompt:

	Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree. The lowest common
	ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants
	(where we allow a node to be a descendant of itself).

	Examples:

	Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1

	    	3
          /   \
         5     1
        / \   / \
       6   2 0   8
          / \
         7   4

	Output: 3

	Explanation: The LCA of nodes 5 and 1 is 3.

	Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4

	Output: 5

	Explanation: The LCA of nodes 5 and 4 is 5, since a node can be a descendant of itself according to the LCA definition.

	Approach:
	- at each node, check:
	  1) is the node itself either p or q?
	  2) if not, perform DFS on children for p and q
	  3) if the node itself is either p or q, do a DFS on the children for the other item
*/
