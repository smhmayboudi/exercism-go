package tree

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
}

func Build(records []Record) (*Node, error) {
	for i := 0; i < len(records); i++ {
		record := records[i]
		nodes := []*Node{}
		if record.ID == record.Parent {
			return &Node{
				ID:       record.ID,
				Children: nodes,
			}, nil
		}
		// nodes = append(nodes, )
	}
	return nil, nil
}
