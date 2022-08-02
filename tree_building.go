package tree

import (
	"errors"
)

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
	root := &Node{ID: -1}
	nodesMap := make(map[int]*Node)
	recordsMap := make(map[int]Record)

	if len(records) == 0 {
		return nil, nil
	}

	maxRecordID := 0
	for _, r := range records {
		if _, ok := recordsMap[r.ID]; ok {
			return nil, errors.New("Found duplicated IDs!")
		} else {
			recordsMap[r.ID] = r
		}

		if maxRecordID < r.ID {
			maxRecordID = r.ID
		}
	}

	for i := 0; i <= maxRecordID; i++ {
		var r Record
		if _, ok := recordsMap[i]; ok {
			r = recordsMap[i]
		} else {
			return nil, errors.New("Tree can not be non-continuous!")
		}

		if r.ID != 0 && r.ID <= r.Parent {
			return nil, errors.New("Parent ID can not be higher than Child ID!")
		}

		newNode := Node{ID: r.ID, Children: []*Node{}}

		if n, ok := nodesMap[r.ID]; ok {
			if n.ID != -1 {
				return nil, errors.New("Found duplicated IDs!")
			} else {
				n.ID = r.ID
				nodesMap[r.ID] = n
			}
		} else {
			nodesMap[r.ID] = &newNode
		}

		if r.ID > 0 {
			if n, ok := nodesMap[r.Parent]; ok {
				n.Children = append(n.Children, &newNode)
				nodesMap[r.Parent] = n
			} else {
				nodesMap[r.Parent] = &Node{ID: -1, Children: []*Node{}}
				n = nodesMap[r.Parent]
				n.Children = append(n.Children, &newNode)
				nodesMap[r.Parent] = n
			}
		} else if r.ID < 0 {
			return nil, errors.New("Record id can not be less than 0!")
		} else {
			if r.Parent > 0 {
				return nil, errors.New("Root node can not have any parent other than itself!")
			}
		}
	}

	if n, ok := nodesMap[0]; ok {
		if n.ID < 0 {
			return nil, errors.New("Root can not be null!")
		} else {
			root = n
		}
	} else {
		return nil, errors.New("Root can not be null!")
	}

	return root, nil
}
