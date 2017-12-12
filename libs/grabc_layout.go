package libs

var Grabc_layout = `
<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
		<title>权限管理</title>
		<link rel="stylesheet" href="/static/bower_extensions/bootstrap/dist/css/bootstrap.min.css?v={{.siteStaticVersion}}" type="text/css">
		<link rel="stylesheet" href="/static/bower_extensions/font-awesome/css/font-awesome.min.css?v={{.siteStaticVersion}}" type="text/css">
		<link rel="stylesheet" href="/static/bower_extensions/admin-lte/dist/css/AdminLTE.min.css?v={{.siteStaticVersion}}" type="text/css">
		<link rel="stylesheet" href="/static/bower_extensions/admin-lte/dist/css/skins/skin-blue.css?v={{.siteStaticVersion}}" type="text/css">
		<link rel="stylesheet" href="/static/css/common.css?v={{.siteStaticVersion}}" type="text/css">
		<script src="/static/bower_extensions/jquery/dist/jquery.min.js?v={{.siteStaticVersion}}"></script>
		<script src="/static/bower_extensions/bootstrap/dist/js/bootstrap.min.js?v={{.siteStaticVersion}}"></script>
		<script src="/static/bower_extensions/admin-lte/dist/js/adminlte.min.js?v={{.siteStaticVersion}}"></script>
    </head>
    <body class="skin-blue sidebar-mini">
        <div class="wrapper">
            <header class="main-header">
                <a href="/site/index" class="logo">
                    {{.app_name}}示例
                </a>

            <nav class="navbar navbar-static-top">
                <a href="#" class="sidebar-toggle" data-toggle="push-menu" role="button">
                    <span class="sr-only">Toggle navigation</span>
                </a>
              <div class="navbar-custom-menu">
                <ul class="nav navbar-nav">  
                  <li class="dropdown user user-menu">
                    <a href="#" class="dropdown-toggle" data-toggle="dropdown">
                      <i class="fa fa-user"></i>
                      <span class="hidden-xs">你好，管理员{{.manage_name}}</span>
                    </a>
                    <ul class="dropdown-menu">
                      <li class="user-footer">
                        <div class="pull-left">
                          <a href="/user/modifypassword" class="btn btn-primary btn-flat">修改密码</a>
                        </div>
                        <div class="pull-right">
                          <a href="/site/logout" class="btn btn-default btn-flat">退出</a>
                        </div>
                      </li>
                    </ul>
                  </li>
                </ul>
              </div>

            </nav>
            </header>
            <aside class="main-sidebar">
            <!-- sidebar: style can be found in sidebar.less -->
            <section class="sidebar" style="height: auto;">
              <!-- sidebar menu: : style can be found in sidebar.less -->
              <ul class="sidebar-menu tree" data-widget="tree">
                <li><a href='/route/index'><i class="fa fa-book"></i> <span>路由管理</span></a></li>
                <li><a href='/permission/index'><i class="fa fa-book"></i> <span>权限管理</span></a></li>
                <li><a href='/role/index'><i class="fa fa-book"></i> <span>角色管理</span></a></li>
                <li><a href='/assignment/index'><i class="fa fa-book"></i> <span>用户授权</span></a></li>
                <li><a href='/menu/index'><i class="fa fa-book"></i> <span>菜单管理</span></a></li>
              </ul>
            </section>
            <!-- /.sidebar -->
          </aside>
            <div class="content-wrapper" style="padding:50px">
              {{.grabc_content}}
            </div>
        </div>
    </body>
</html>
`
