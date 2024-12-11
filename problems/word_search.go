package problems

func dfs(board [][]byte, i, j int, node *Trie, visited [][]bool, str string, ans *map[string]bool) {
    m, n := len(board), len(board[0])

    // Base case: out of bounds, already visited, or no valid Trie child
    if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] || node.child[board[i][j]-'a'] == nil {
        return
    }

    // Traverse the Trie and current board cell
    char := board[i][j]
    node = node.child[char-'a']
    str += string(char)

    // If the word ends at this node, add it to the answer set
    if node.wordEnd {
        (*ans)[str] = true
    }

    // Mark the cell as visited
    visited[i][j] = true

    // Explore all 4 directions
    directions := [][]int{
        {1, 0},  // Down
        {0, 1},  // Right
        {-1, 0}, // Up
        {0, -1}, // Left
    }

    for _, dir := range directions {
        dfs(board, i+dir[0], j+dir[1], node, visited, str, ans)
    }

    // Backtrack
    visited[i][j] = false
}

func findWords(board [][]byte, words []string) []string {
    // Build Trie from the words list
    myTrie := TrieConstructor()
    for _, word := range words {
        myTrie.TrieInsert(word)
    }

    // Initialize visited array
    visited := make([][]bool, len(board))
    for i := range board {
        visited[i] = make([]bool, len(board[0]))
    }

    // Use a map to avoid duplicate words in the answer
    ans := make(map[string]bool)

    // Perform DFS from every cell in the board
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            dfs(board, i, j, &myTrie, visited, "", &ans)
        }
    }

    // Convert the result map to a slice
    result := []string{}
    for word := range ans {
        result = append(result, word)
    }

    return result
}