// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/ngocphuongnb/tetua/packages/entrepository/ent/topic"
)

// Topic is the model entity for the Topic schema.
type Topic struct {
	config `json:"-" validate:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty" validate:"max=255"`
	// Slug holds the value of the "slug" field.
	Slug string `json:"slug,omitempty" validate:"max=255"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty" validate:"max=255"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty" validate:"required"`
	// ContentHTML holds the value of the "content_html" field.
	ContentHTML string `json:"content_html,omitempty" validate:"required"`
	// ParentID holds the value of the "parent_id" field.
	ParentID int `json:"parent_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TopicQuery when eager-loading is set.
	Edges TopicEdges `json:"edges"`
}

// TopicEdges holds the relations/edges for other nodes in the graph.
type TopicEdges struct {
	// Posts holds the value of the posts edge.
	Posts []*Post `json:"posts,omitempty"`
	// Children holds the value of the children edge.
	Children []*Topic `json:"children,omitempty"`
	// Parent holds the value of the parent edge.
	Parent *Topic `json:"parent,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// PostsOrErr returns the Posts value or an error if the edge
// was not loaded in eager-loading.
func (e TopicEdges) PostsOrErr() ([]*Post, error) {
	if e.loadedTypes[0] {
		return e.Posts, nil
	}
	return nil, &NotLoadedError{edge: "posts"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e TopicEdges) ChildrenOrErr() ([]*Topic, error) {
	if e.loadedTypes[1] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TopicEdges) ParentOrErr() (*Topic, error) {
	if e.loadedTypes[2] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: topic.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Topic) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case topic.FieldID, topic.FieldParentID:
			values[i] = new(sql.NullInt64)
		case topic.FieldName, topic.FieldSlug, topic.FieldDescription, topic.FieldContent, topic.FieldContentHTML:
			values[i] = new(sql.NullString)
		case topic.FieldCreatedAt, topic.FieldUpdatedAt, topic.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Topic", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Topic fields.
func (t *Topic) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case topic.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case topic.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case topic.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				t.UpdatedAt = value.Time
			}
		case topic.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				t.DeletedAt = value.Time
			}
		case topic.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = value.String
			}
		case topic.FieldSlug:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field slug", values[i])
			} else if value.Valid {
				t.Slug = value.String
			}
		case topic.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				t.Description = value.String
			}
		case topic.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				t.Content = value.String
			}
		case topic.FieldContentHTML:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content_html", values[i])
			} else if value.Valid {
				t.ContentHTML = value.String
			}
		case topic.FieldParentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field parent_id", values[i])
			} else if value.Valid {
				t.ParentID = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryPosts queries the "posts" edge of the Topic entity.
func (t *Topic) QueryPosts() *PostQuery {
	return (&TopicClient{config: t.config}).QueryPosts(t)
}

// QueryChildren queries the "children" edge of the Topic entity.
func (t *Topic) QueryChildren() *TopicQuery {
	return (&TopicClient{config: t.config}).QueryChildren(t)
}

// QueryParent queries the "parent" edge of the Topic entity.
func (t *Topic) QueryParent() *TopicQuery {
	return (&TopicClient{config: t.config}).QueryParent(t)
}

// Update returns a builder for updating this Topic.
// Note that you need to call Topic.Unwrap() before calling this method if this Topic
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Topic) Update() *TopicUpdateOne {
	return (&TopicClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Topic entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Topic) Unwrap() *Topic {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Topic is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Topic) String() string {
	var builder strings.Builder
	builder.WriteString("Topic(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(t.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", deleted_at=")
	builder.WriteString(t.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", name=")
	builder.WriteString(t.Name)
	builder.WriteString(", slug=")
	builder.WriteString(t.Slug)
	builder.WriteString(", description=")
	builder.WriteString(t.Description)
	builder.WriteString(", content=")
	builder.WriteString(t.Content)
	builder.WriteString(", content_html=")
	builder.WriteString(t.ContentHTML)
	builder.WriteString(", parent_id=")
	builder.WriteString(fmt.Sprintf("%v", t.ParentID))
	builder.WriteByte(')')
	return builder.String()
}

// Topics is a parsable slice of Topic.
type Topics []*Topic

func (t Topics) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}