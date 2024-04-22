package tree

import "fmt"

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
	if len(records) == 0 {
		return nil, nil
	}

	const rootID = 0
	positions := make([]int, len(records))
	for i := 0; i < len(records); i++ {
		record := records[i]
		if record.ID < rootID || record.ID >= len(records) {
			return nil, fmt.Errorf("out of bounds record id %d", record.ID)
		}
		positions[record.ID] = i
	}

	nodes := make([]Node, len(records))
	for i := 0; i < len(positions); i++ {
		r := records[positions[i]]
		if r.ID != i {
			return nil, fmt.Errorf("non-contiguous node %d (want %d)", r.ID, i)
		}
		validParentForChild := (r.ID > r.Parent) || (r.ID == rootID && r.Parent == rootID)
		if !validParentForChild {
			return nil, fmt.Errorf("node %d has impossible parent %d", r.ID, r.Parent)
		}
		nodes[i].ID = i
		if i != rootID {
			p := &nodes[r.Parent]
			p.Children = append(p.Children, &nodes[i])
		}
	}
	return &nodes[0], nil
}
