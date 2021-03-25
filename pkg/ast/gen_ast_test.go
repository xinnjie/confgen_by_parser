package ast

//func TestGenAst(t *testing.T) {
//	testString := "BarStruct"
//	type args struct {
//		defineString string
//	}
//	tests := []struct {
//		name string
//		args args
//		want Map
//	}{
//		{
//			"one Key map parser test ",
//			args{defineString: "SingleKeyMap\n" +
//				"foo_key  uint32 \"这是scalar字段\" \n" +
//				"foo_uint32  uint32 \"\" \n" +
//				"bar vector BarStruct \"这是vector<struct>字段\" \n" +
//				"{bar_1  uint32 \"\" \n" +
//				"bar_2  uint32 \"\" }\n" +
//				"$\n" +
//				"{BarEnum} \n" +
//				"[实例类型]实例类型1 1 E_BAR_1 \n" +
//				"[实例类型]实例类型2 2 E_BAR_2 \n",
//			},
//			Map{
//				ID: "SingleKeyMap",
//				Fields: []*Field{
//					{
//						Name: "foo_key",
//						Declare: TypeDeclare{
//							Scalar:     &Scalar{IsUINT32: true},
//							StructName: nil,
//							Vector:     nil,
//						},
//						Desc:       `"这是scalar字段"`,
//						StructList: nil,
//					},
//					{
//						Name: "foo_uint32",
//						Declare: TypeDeclare{
//							Scalar:     &Scalar{IsUINT32: true},
//							StructName: nil,
//							Vector:     nil,
//						},
//						Desc:       `""`,
//						StructList: nil,
//					},
//					{
//						Name: "bar",
//						Declare: TypeDeclare{
//							Scalar: nil,
//							Vector: &VectorType{
//								ElementType: &TypeDeclare{
//									StructName: &testString,
//								},
//							},
//						},
//						Desc: `"这是vector<struct>字段"`,
//						StructList: []*Struct{{Fields: []*StructElement{
//							{
//								Id:   "bar_1",
//								Type: &Scalar{IsUINT32: true},
//								Desc: `""`,
//							},
//							{
//								Id:   "bar_2",
//								Type: &Scalar{IsUINT32: true},
//								Desc: `""`,
//							},
//						}}},
//					},
//				},
//				Enums: []*Enum{
//					{
//						EnumType: "BarEnum",
//						EnumElms: []*EnumElement{
//							{
//								EnumLiteral: "[实例类型]实例类型1",
//								EnumValue:   1,
//								ID:          "E_BAR_1",
//							},
//							{
//								EnumLiteral: "[实例类型]实例类型2",
//								EnumValue:   2,
//								ID:          "E_BAR_2",
//							},
//						},
//					},
//				},
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			defineString := bytes.NewBufferString(tt.args.defineString)
//
//			got := GenAst(defineString, tt.name)
//			assert.EqualValues(t, tt.want, got)
//		})
//	}
//}
