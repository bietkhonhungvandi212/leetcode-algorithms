# Intuition

To find the lowest common ancestor (LCA) of two given nodes in a binary tree, we need to identify the first node that is an ancestor of both. This can be achieved by tracing the paths from each node to the root and finding their intersection. To track these paths efficiently, I use a hash map to store parent-child relationships for each node.

# Approach
1. Use a hash map to store parent-child relationships, where the key is a child node (left or right) and the value is its parent node:
    - If node.left is not null, set mapper[node.left] = node.
    - If node.right is not null, set mapper[node.right] = node.

2. Initialize a set to store all ancestor nodes of the first given node, q.
3. Traverse from node q to the root using the hash map, adding each parent node to the set.
4. Traverse from the second given node, p, to the root, checking each parent node against the set. The first parent node found in the set is the lowest common ancestor.