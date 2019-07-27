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
    <div class="x-nav">
      <span class="layui-breadcrumb">
        <a href="/">首页</a>
        <a href="javascript:location.replace(location.href);">管理员管理</a>
        <a>
          <cite>管理员列表</cite></a>
      </span>
      <a class="layui-btn layui-btn-small" style="line-height:1.6em;margin-top:3px;float:right" href="javascript:location.replace(location.href);" title="刷新">
        <i class="layui-icon" style="line-height:30px">ဂ</i></a>
    </div>
    <div class="x-body">
      <div class="layui-row">
        <form class="layui-form layui-col-md12 x-so">
          <input type="text" name="username"  placeholder="请输入用户名" autocomplete="off" class="layui-input">
          <button class="layui-btn"  lay-submit="" lay-filter="sreach"><i class="layui-icon">&#xe615;</i></button>
        </form>
      </div>
      <xblock>
          <!--<button class="layui-btn layui-btn-danger" onclick="delAll()"><i class="layui-icon"></i>批量删除</button> -->
          {{ if $.addName}}
        <button class="layui-btn" onclick="x_admin_show('添加用户','/{{$.addAction}}')"><i class="layui-icon"></i>{{$.addName}}</button>
          {{end}}

        <span class="x-right" style="line-height:40px">共有数据：{{$.total}} 条</span>
      </xblock>
      <table class="layui-table">
        <thead>
          <tr>
              <!--<th>
                 <div class="layui-unselect header layui-form-checkbox" lay-skin="primary"><i class="layui-icon">&#xe605;</i></div>
            </th>-->
            <th>ID</th>
            <th>管理员名</th>
            <th>角色</th>
            <th>创建时间</th>
            <th width="200">操作</th>
        </thead>
        <tbody>
        {{range $key, $val := $.admins}}
          <tr>
            <!-- <td>
             <div class="layui-unselect layui-form-checkbox" lay-skin="primary" data-id='{{$val.Id}}'><i class="layui-icon">&#xe605;</i></div>
            </td> -->
            <td>{{$val.Id}}</td>
            <td>{{$val.Username}}</td>
            <td>{{$val.Role}}</td>
            <td> {{$val.Create_time | timeFormat }}</td>

            <td class="td-manage">
                {{if $.editName}}
              <a title="编辑"  onclick="x_admin_show('{{$.editName}}','/{{$.editAction}}?id={{$val.Id}}')" href="javascript:;">
                <span class="layui-btn">修改&nbsp;  </span>
              </a>
              {{end}}
               {{if $.removeName}}
              <a title="删除" onclick="member_del(this,'/{{$.removeAction}}?id={{$val.Id}}')" href="javascript:;">
                <span class="layui-btn" style="margin-left: 10px; background-color: #993333">&nbsp;  删除</span>
              </a>
               {{end}}

            </td>
          </tr>
        {{end}}
        </tbody>
      </table>
      <div class="page">
        <div>

            <a class="prev" href="{{$.prePage}}">上一页</a>

            <span class="current"> 当前第  {{$.pageInfo}}  页   </span>

            <span class="current">  共  {{$.totalPage}}  页</span>
            <a class="next" href="{{$.nextPage}}">下一页</a>


        </div>
      </div>

    </div>
    <script>
      layui.use('laydate', function(){
        var laydate = layui.laydate;
        
        //执行一个laydate实例
        laydate.render({
          elem: '#start' //指定元素
        });

        //执行一个laydate实例
        laydate.render({
          elem: '#end' //指定元素
        });
      });


      /*用户-删除*/
      function member_del(obj,url){
          layer.confirm('确认要删除吗？',function(index){
              $.get(url, function(res){
                  if(!res){
                      layer.msg('系统异常，请联系技术员!',{icon:2,time:1000});
                      return false;
                  }
                  if(res.code === 200){
                      $(obj).parents("tr").remove();
                      layer.msg('已删除!',{icon:1,time:1000});
                  }else{
                      layer.msg(res.msg,{icon:2,time:1000});
                  }
              });
          });
      }

    </script>
    <script>var _hmt = _hmt || []; (function() {
        var hm = document.createElement("script");
        hm.src = "https://hm.baidu.com/hm.js?b393d153aeb26b46e9431fabaf0f6190";
        var s = document.getElementsByTagName("script")[0];
        s.parentNode.insertBefore(hm, s);
      })();</script>
  </body>

</html>