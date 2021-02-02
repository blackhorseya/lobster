package todo

import "testing"

func TestTask_ToLineByFormat(t1 *testing.T) {
	type fields struct {
		ID        string
		Title     string
		Completed bool
		CreateAt  int64
	}
	type args struct {
		format string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "format then string",
			fields: fields{
				ID:        "4cc9a0f9-66ac-4ed8-b6fd-5afbf26e4453",
				Title:     "task1",
				Completed: true,
				CreateAt:  1612223933070772764,
			},
			args: args{format: "%-36s\t%-20s\t%-6v\t%-9v"},
			want: "4cc9a0f9-66ac-4ed8-b6fd-5afbf26e4453\ttask1               \ttrue  \t2021-02-02 07:58:53.070772764 +0800 CST",
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Task{
				ID:        tt.fields.ID,
				Title:     tt.fields.Title,
				Completed: tt.fields.Completed,
				CreateAt:  tt.fields.CreateAt,
			}
			if got := t.ToLineByFormat(tt.args.format); got != tt.want {
				t1.Errorf("ToLineByFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}
