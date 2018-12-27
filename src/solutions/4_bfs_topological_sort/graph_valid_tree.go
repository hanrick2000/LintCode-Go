func validTree(n int, edges [][]int) bool {
       if len(edges) != n - 1 {
           return false
       }
       nextNodes := make([][]int, n)
       for _, edge := range edges {
           nextNodes[edge[0]] = append(nextNodes[edge[0]], edge[1])
           nextNodes[edge[1]] = append(nextNodes[edge[1]], edge[0])
       }
      visited := make(map[int]int)
      for i := 0; i < n; i++ {
          visited[i] = 0
      }
      for i := 0; i < n; i++ {
          if !dfs(-1, i, visited, nextNodes) {
              return false
          }
      }
      return true
  }
  
 func dfs(prev int, cur int, visited map[int]int, nextNodes [][]int) bool {
      val, _ := visited[cur]
      if val == 1 {
          return true
      }
      if val == 2 && prev != val {
          return false
      }
      visited[cur] = 2
      for _, next := range nextNodes[cur] {
          if next == prev {
              continue
          }
          if !dfs(cur, next, visited, nextNodes) {
              return false
          }
      }
      visited[cur] = 1
      return true
 }

