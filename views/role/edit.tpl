<!DOCTYPE html>
<html>
  
  <head>
    <meta charset="UTF-8">
    <title>欢迎页面-X-admin2.0</title>
    <meta name="renderer" content="webkit">
    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
    <meta name="viewport" content="width=device-width,user-scalable=yes, minimum-scale=0.4, initial-scale=0.8,target-densitydpi=low-dpi" />
    <link rel="shortcut icon" href="/static/favicon.ico" type="image/x-icon" />
    <link rel="stylesheet" href="/static/css/font.css">
    <link rel="stylesheet" href="/static/css/xadmin.css">
    <script type="text/javascript" src="https://cdn.bootcss.com/jquery/3.2.1/jquery.min.js"></script>
    <script type="text/javascript" src="/static/lib/layui/layui.js" charset="utf-8"></script>
    <script type="text/javascript" src="/static/js/xadmin.js"></script>
    <!-- 让IE8/9支持媒体查询，从而兼容栅格 -->
    <!--[if lt IE 9]>
      <script src="https://cdn.staticfile.org/html5shiv/r29/html5.min.js"></script>
      <script src="https://cdn.staticfile.org/respond.js/1.4.2/respond.min.js"></script>
    <![endif]-->
  </head>
  
  <body>
    <div class="x-body">
        <form class="layui-form">
            <div class="layui-form-item">
                <label for="name" class="layui-form-label">
                    角色名称
                </label>
                <div class="layui-input-inline">
                    <input type="text" id="name" name="name" required="" value="{{.role.Name}}" autocomplete="off" class="layui-input">
                </div>
            </div>
            <div class="layui-form-item layui-form-text">
                <label class="layui-form-label">
                    权限列表
                </label>
                <table  class="layui-table layui-input-block">
                    <tbody>

                    {{range $ke,$vo := $.nav}}
                        <tr>
                            <td>
                                <input name="ids" type="checkbox"  value="{{$vo.id}}" title="{{$vo.name}}" lay-filter='group' id='group_{{$vo.id}}' {{if Checked $vo.id $.role.Minu_id }} checked="checked"  {{end}}>
                            </td>
                            <td>
                                <div class="layui-input-block">
                                   {{range $key,$val := $.menu}}
                                    {{if eq $val.mid $vo.id}}
                                        <input name="ids" type="checkbox"  value="{{$val.id}}" title="{{$val.name}}" class='group_{{$vo.id}}' lay-filter='group_child' check_child='group_{{$vo.id}}' {{if Checked $val.id $.role.Minu_id }} checked="checked"  {{end}}>
                                    {{end}}
                                   {{end}}
                                    {{range $k,$v :=$.action}}
                                    {{if eq $v.mid $vo.id}}
                                    <input name="ids" type="checkbox"  value="{{$v.id}}" title="{{$v.name}}" class='group_{{$vo.id}}' lay-filter='group_child' check_child='group_{{$vo.id}}' {{if Checked $v.id $.role.Minu_id }} checked="checked"  {{end}}>
                                    {{end}}
                                    {{end}}
                                </div>

                            </td>
                        </tr>
                    {{end}}
                    </tbody>
                </table>
            </div>
            <div class="layui-form-item">
                <input type="hidden"  name="role_id" value='{{.role.Id}}' >
                <button  class="layui-btn" lay-filter="add" lay-submit="" id='btn'>立即提交</button>
            </div>
      </form>
    </div>
    <script>

        layui.use(['form','layer'], function(){
            $ = layui.jquery;
            var form = layui.form
                ,layer = layui.layer;
            //父节点全选/全不选
            form.on('checkbox(group)', function(data){
                var check = "."+data.elem.id+"";//获取父节点的id
                $(check).prop('checked',$(this).prop("checked"));//将和父节点id相同class的子节点选中
                form.render();//更新全部
            });

            //子节点
            form.on('checkbox(group_child)',function(data){
                var check = '.'+data.elem.attributes['check_child'].value;//当前子节点下的check_child的值
                var idcheck = '#'+data.elem.attributes['check_child'].value;
                var test = 1;//定义一个初始值
                for (var i = $(check).length - 1; i >= 0; i--) {
                    if($(check).eq(i).prop("checked")){
                        test = 2;//当有复选框被选中时，初始值更改为2
                    }
                }

                if(test == 1){
                    $(idcheck).prop('checked',false);//如果初始值依旧为一，说明没有子节点的复选框被选中,则将和子节点相同class的
                }else{
                    $(idcheck).prop('checked',true);//否则将父节点全选框选中
                }
                form.render();//更新全部
            });
            //监听提交
            form.on('submit(add)', function(data){
                var arr = new Array();


                //获取复选框所有选中的值
                $("input:checkbox[name='ids']:checked").each(function(i){
                    arr[i] = $(this).val();
                });
                data.field.ids = arr.join(",");//将数组合并成字符串
                $.post("{{.act}}",data.field,function(res){
                    if(!res){
                        layer.msg('系统异常，请联系技术员!',{icon:2,time:1000});
                        return false;
                    }
                    if(res.code === 200){
                        layer.alert("修改成功", {icon: 6},function () {
                            window.parent.location.reload();
                            var index = parent.layer.getFrameIndex(window.name);
                            parent.layer.close(index);
                        });
                    }else{
                        layer.msg(res.msg,{icon:2,time:2000});
                    }
                });
                return false;
            });


        });
    </script>
    <script>var _hmt = _hmt || []; (function() {
        var hm = document.createElement("script");
        hm.src = "https://hm.baidu.com/hm.js?b393d153aeb26b46e9431fabaf0f6190";
        var s = document.getElementsByTagName("script")[0];
        s.parentNode.insertBefore(hm, s);
      })();</script>
  </body>

</html>