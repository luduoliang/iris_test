layui.use(['element', 'layer'], function() {
    var element = layui.element(),
        $ = layui.jquery,
        layer = layui.layer;

    //iframe自适应
    $(window).on('resize', function() {
        var $content = $('.admin-nav-card .layui-tab-content');
        $content.height($(this).height() - 147);
        $content.find('iframe').each(function() {
            $(this).height($content.height());
        });
    }).resize();

    //添加tab
    var $tabs = $('#admin-tab');
    var $container = $('#admin-tab-container');
    //绑定 nav 点击事件
    $('ul.admin-nav-tree').find('dd > a').each(function() {
        var $this = $(this);
        //获取设定的url
        var url = $this.data('url');
        if(url !== undefined) {
            $this.on('click', function() {
                var aHtml = $this.html(),
                    title = $(this).children('cite').text();
                addTab('admin-tab', aHtml, url, title, false);
            });
        }
    });
    //添加新的iframe
    $('.add_wrapper').on('click', function () {
        var url = $(this).data('url'),
            title = $(this).data('title'),
            icon = $(this).data('icon');
        icon = (undefined == icon) ? 'dot-circle-o' : icon;
        if (undefined !== url && undefined !== title) {
            var aHtml = '<i class="fa fa-' + icon + '" aria-hidden="true"></i><cite>' + title + '</cite>';
            addTab('admin-tab',aHtml, url, title, true);
        }
    });

    $('#user').on('click', function() {
        $('#user-item').toggle();
    });

    //刷新框架中的内容
    $('.refresh_wrapper').on('click', function(){
        var $current_iframe=$("#admin-tab-container div:visible iframe");
        $current_iframe[0].contentWindow.location.reload();
        return false;
    });

    //关闭当前框架
    $('.close_wrapper').on('click', function(){
        var tab_close = $("#admin-tab .layui-this .layui-tab-close", parent.document);
        if(tab_close.length === 1){
            var iframes = $("#admin-tab-container div iframe", parent.document);
            if(iframes.length > 1) iframes[iframes.length-2].contentWindow.location.reload();
            tab_close.click();
        }
        return false;
    });

    /**
     * 添加tab
     * @param string toTab 添加tab的目标
     * @param string title 标题Html内容
     * @param string url iframe内容地址
     * @param string name 标题名，用于区分是否已经添加
     * @param boolean isChild 是否来自于子窗口的调用
     * @returns {boolean}
     */
    function addTab(toTab, title, url, name, isChild) {
        var count = 0,
            iframe = '<iframe src="' + url + '"></iframe>',
            tabIndex;
        if (isChild) {
            var $Tabs = $('#' + toTab, window.parent.document),
                element = window.parent.layui.element();
        } else {
            var $Tabs = $('#' + toTab),
                element = layui.element();
        }
        
        if (undefined == $Tabs) {
            return true;
        }
       
        $Tabs.find('li').each(function (i) {
            if (name === $(this).find('cite').text()) {
                count++;
                tabIndex = i;
            }
        });
        //tab不存在
        if (count === 0) {
            //添加删除样式
            title += '<i class="layui-icon layui-unselect layui-tab-close">&#x1006;</i>';
            //添加tab
            element.tabAdd(toTab, {
                title: title,
                content: iframe
            });
            //iframe 自适应
            var $content = $Tabs.next('div');
            $content.find('iframe').each(function () {
                $(this).height($content.height());
            });

            var $li = $Tabs.find('li');

            $li.eq($li.length - 1).children('i.layui-tab-close').on('click', function () {
                element.tabDelete(toTab, $(this).parent('li').index()).init();
            });
            //获取焦点
            element.tabChange(toTab, $li.length - 1);
        } else {
            element.tabChange(toTab, tabIndex);
        }
    }

    //手机设备的简单适配
    var treeMobile = $('.site-tree-mobile'),
        shadeMobile = $('.site-mobile-shade');
    treeMobile.on('click', function() {
        $('body').addClass('site-mobile');
    });
    shadeMobile.on('click', function() {
        $('body').removeClass('site-mobile');
    });
});