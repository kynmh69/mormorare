package time

import "testing"

func TestDateTime_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		d       DateTime
		args    args
		wantErr bool
	}{
		{
			name: "TestDateTime_UnmarshalJSON",
			d:    DateTime{},
			args: args{
				data: []byte(`"2020-01-01"`),
			},
			wantErr: false,
		},
		{
			name: "TestDateTime_UnmarshalJSON",
			d:    DateTime{},
			args: args{
				data: []byte(`"2020-01-01 00:00:00"`),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
