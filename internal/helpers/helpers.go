package helpers

import (
	"ringhover-go/internal/domain/models"
	"ringhover-go/internal/domain/resp"
)

// BuildSubtasksForest builds a forest (array of trees) composed of the
// direct children of the given rootID, each with its own descendants.
// Input is a flat slice of tasks that are known to be descendants of rootID
// (but does NOT include the root task itself).
func BuildSubtasksForest(rows []models.Task, rootID uint64) []resp.TaskTree {
	// Create all nodes and index them by ID for O(1) lookups later.
	nodes := make(map[uint64]*resp.TaskTree, len(rows))
	for i := range rows {
		nodes[rows[i].Id] = &resp.TaskTree{
			Task: resp.Task{
				Id:           rows[i].Id,
				Title:        rows[i].Title,
				Description:  rows[i].Description,
				Status:       rows[i].Status,
				Priority:     rows[i].Priority,
				DueDate:      rows[i].DueDate,
				CompletedAt:  rows[i].CompletedAt,
				ParentTaskID: rows[i].ParentTaskID,
				CategoryID:   rows[i].CategoryID,
				CreatedAt:    rows[i].CreatedAt,
				UpdatedAt:    rows[i].UpdatedAt,
			},
		}
	}

	//  Link parent → children for nodes that are present in the 'rows' set.
	for _, n := range nodes {
		if n.ParentTaskID != nil {
			if p := nodes[*n.ParentTaskID]; p != nil {
				p.Children = append(p.Children, *n)
			}
		}
	}

	// The forest roots are the tasks whose direct parent is 'rootID'
	var roots []resp.TaskTree
	for _, n := range nodes {
		if n.ParentTaskID != nil && *n.ParentTaskID == rootID {
			roots = append(roots, *n)
		}
	}
	return roots
}
