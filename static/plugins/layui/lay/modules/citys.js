/**
 * layui-v1.0.7 省市三级联动modules
 * dingweixian@imohoo.com
 * 2016-12-22
 * @use
 * pca_options = {
 *     url : '{:U("Index/getCityList")}',
 *     lang : 'cn', //语言cn-中文;en-英文
 *     country : 'select[name=country]',
 *     province : 'select[name=province]',
 *     city : 'select[name=city]',
 *     initCity : '北京',
 * };
 */
layui.define('layer', function(exports){
    "use strict";

    var $ = layui.jquery;
    var pca = {};

    pca.keys = {};
    pca.ckeys = {};
    pca.options = {
        url : '', //服务器地址
        lang : 'cn', //语言 cn-中文[默认];en-英文
        country : '', //select[name=country]
        province : '', //select[name=province]
        city : '', //select[name=city]
        area : '', //select[name=area]
        initCountry : '',
        initProvince : '',
        initCity : '',
        initArea : ''
    };

    pca.initSetting = function (options) {
        for (var i in options) {
            if (pca.options.hasOwnProperty(i)) {
                pca.options[i] = options[i];
            }
        }
    };

    pca.init = function(options){
        pca.initSetting(options);
        if(pca.options.url == '') return;
        var country = options.country,
            province = options.province,
            city = options.city,
            area = options.area;
        if(!country || !$(country).length){
            pca.getData(1,2); //默认中国
        }else{ //三级联动中有国家
            pca.getData(0,0); //获取国家数据
        }

        if(!province || !$(province).length) return;
        pca.formRender(province);
        layui.form().on('select(country)', function(data){
            pca.getData(data.value, 2); //切换后获取城市数据
            pca.formHidden('country', data.value);
        });

        if(!city || !$(city).length) return;
        pca.formRender(city);
        layui.form().on('select(province)', function(data){
            pca.getData(data.value, 3);
            pca.formHidden('province', data.value);
        });

        if(!area || !$(area).length) return;
        pca.formRender(area);
        layui.form().on('select(city)', function(data){
            pca.getData(data.value, 4);
            pca.formHidden('city', data.value);
        });
        layui.form().on('select(area)', function(data){
            pca.formHidden('area', data.value);
        });
    };

    //请求服务端获取信息
    pca.getData = function(pid, level){
        $.post(pca.options.url,{pid:pid,level:level},function(data){
            if(data.status == 'success'){
                pca.selectAppend(data.data);
            }
        })
    };

    pca.formRender = function(obj){
        $(obj).html('');
        $(obj).append('<option></option>');
        layui.form('select').render();
    };

    //扩展城市信息
    pca.selectAppend = function(data){
        var select, level = parseInt(data[0].level), autoload = true, initSelect = null;
        switch (level){
            case 1: //国家
                select = pca.options.country;
                break;
            case 2: //省份
                select = pca.options.province;
                break;
            case 3: //市
                select = pca.options.city;
                break;
            case 4: //县
                select = pca.options.area;
                break;
            default: //未获取到值
                return;
        }
        $(select).html('');
        $(select).append('<option selected>' + ((pca.options.lang == "en") ? "All" : "全部") + '</option>');
        if(!select || !$(select).length) return;

        for(var i in data){
            if(pca.options.lang == 'en'){
                $(select).append('<option value="' + data[i].id + '">' + data[i].ename + '</option>');
                pca.keys[data[i].ename] = data[i];
            }else{
                $(select).append('<option value="' + data[i].id + '">' + data[i].cname + '</option>');
                pca.keys[data[i].cnname] = data[i];
            }
        }
        $(select).find('option:eq(1)').attr('selected', true); //第一个显示出来
        layui.form('select').render(); //重新渲染表单
        switch (level){
            case 1:
                if(pca.options.initCountry) initSelect = $(pca.options.country).next().find('dd:contains("' + pca.options.initCountry + '")');
                break;
            case 2:
                if(pca.options.initProvince) initSelect = $(pca.options.province).next().find('dd:contains("' + pca.options.initProvince + '")');
                break;
            case 3:
                if(pca.options.initCity) initSelect = $(pca.options.city).next().find('dd:contains("' + pca.options.initCity + '")');
                break;
            case 4:
                if(pca.options.initArea) initSelect = $(pca.options.area).next().find('dd:contains("' + pca.options.initArea + '")');
                break;
        }
        if (initSelect && initSelect.length > 0) {
            autoload = false;
            initSelect.click();
        }
        if(autoload) $(select).next().find('.layui-this').removeClass('layui-this').click(); //自动加载下一级
    };

    pca.formHidden = function(obj, val){
        if(!$('#pca-hide-'+obj).length)
            $('body').append('<input id="pca-hide-'+obj+'" type="hidden" value="'+val+'" />');
        else
            $('#pca-hide-'+obj).val(val);
    };

    //暴露接口
    exports('pca', function(options){
        pca.init(options);
    });
});