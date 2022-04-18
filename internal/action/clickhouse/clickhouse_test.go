package clickhouse

import (
	"context"
	"fmt"
	"testing"
)

func TestTableList(t *testing.T) {
	ctx := context.Background()
	tp := NewTransport()
	elem, err := tp.TableList(ctx,
		[]string{
			"http://default:qingcloud2019@139.198.18.173:8123/",
		})
	if err != nil {
		t.Errorf("Check() error = %v", err)
	}
	for _, tb := range elem {
		t.Log("table:", tb)
	}
}
func TestTransport_TableList(t *testing.T) {
	type args struct {
		Endpoints []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Error",
			args{
				[]string{
					"http://default:mdmp2019@139.198.186.46:18123",
				},
			},
			false,
		},
		{
			"Error",
			args{
				[]string{
					"http://default:mdmp2019@139.198.186.46:18123?database=default",
				},
			},
			true,
		},
	}
	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Transport{}
			_, err := this.TableList(ctx, tt.args.Endpoints)
			if (err != nil) != tt.wantErr {
				t.Errorf("TableInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(gotElem, tt.wantElem) {
			//	t.Errorf("TableInfo() gotElem = %v, want %v", gotElem, tt.wantElem)
			//}
		})
	}
}

func TestTableInfo(t *testing.T) {
	ctx := context.Background()
	tp := NewTransport()
	elem, err := tp.TableInfo(ctx,
		[]string{
			"http://default:mdmp2019@139.198.186.46:18123/",
		},
		"table_test")
	if nil != err {
		fmt.Println(err)
		return
	}
	for _, field := range elem.Fields {
		fmt.Println(field)
	}

	fmt.Println(elem.PartitionFields)
	fmt.Println(elem.OrderByFields)

	if err != nil {
		t.Errorf("Check() error = %v", err)
	}
}

func TestExceSQL(t *testing.T) {
	ctx := context.Background()
	tp := NewTransport()
	err := tp.ExceSQL(ctx,
		[]string{
			"http://default:mdmp2019@139.198.186.46:18123/",
		},
		//"CREATE TABLE default.create_table_test21(aaa Float64) ENGINE = MergeTree()  PARTITION BY (aaa)  ORDER BY (aaa) SETTINGS index_granularity = 8192")
		"CREATE TABLE default.create_table_test_011(aaa Float32,bbb Float32) ENGINE = MergeTree()  PARTITION BY (aaa)  ORDER BY (bbb) SETTINGS index_granularity = 8192")
	fmt.Println(err)

	if err != nil {
		t.Errorf("Check() error = %v", err)
	}
}
