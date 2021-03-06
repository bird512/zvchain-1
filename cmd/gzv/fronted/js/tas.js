
layui.use(['form', 'jquery', 'element', 'layer', 'table'], function(){
    var element = layui.element;
    var form = layui.form;
    var layer = layui.layer;
    var $ = layui.$;
    var HOST = "/";
    var ref;
    var host_ele = $("#host");
    var online=false;
    var current_block_height = -1;
    var last_sync_block = -1;

    var current_group_height = -1;
    var last_sync_group = -1;

    host_ele.text(HOST);
    var blocks = [];
    var groups = [];


    var lastReloadBlockSize = 0;
    var lastReloadGroupSize = 0;
    var workGroups = [];
    var groupIds = new Set();
    var table = layui.table;
    var dashboard_update_switch = true;


    var block_table = table.render({
        elem: '#block_detail' //指定原始表格元素选择器（推荐id选择器）
        ,initSort: {
            field: 'height',
            type: 'desc'
        }
        ,cols: [[{field:'height',title: '块高', sort:true},
            {field:'hash', title: 'hash', templet: '<div><a href="javascript:void(0);" class="layui-table-link" name="block_table_hash_row">{{d.hash}}</a></div>'},
            {field:'pre_hash', title: 'pre_hash'},{field:'pre_time', title: 'pre_time', width: 189},{field:'cur_time', title: 'cur_time', width: 189},
            {field:'castor', title: 'castor'},{field:'group_id', title: 'group_id'}, {field:'txs', title: 'tx_count'}, {field:'qn', title: 'qn'}
            , {field:'total_qn', title: 'totalQN'}]] //设置表头
        ,data: blocks
        ,page: false
        ,limit:200
    });

    var group_table = table.render({
        elem: '#group_detail' //指定原始表格元素选择器（推荐id选择器）
        ,initSort: {
            field: 'begin_height',
            type: 'desc'
        }
        ,cols: [[{field:'id',title: '组seed', width:140}, {field:'threshold',title: '门限', width:80},
            {field:'begin_height', title: '生效高度', width: 100},{field:'dismiss_height', title: '解散高度', width:100},{field:'mem_size', title: '成员数量', width:100},
            {field:'members', title: '成员列表'}]] //设置表头
        ,data: groups
        ,page: false
        ,limit:200
    });

    var work_group_table = table.render({
        elem: '#work_group_detail' //指定原始表格元素选择器（推荐id选择器）
        ,initSort: {
            field: 'begin_height',
            type: 'desc'
        }
        ,cols: [[{field:'id',title: '组seed', width:140}, {field:'threshold',title: '门限', width:80},
            {field:'begin_height', title: '生效高度', width: 100},{field:'dismiss_height', title: '解散高度', width:100},{field:'mem_size', title: '成员数量', width:100},
            {field:'members', title: '成员列表'}]] //设置表头
        ,data: groups
        ,page: false
        ,limit:200
    });

    let reward_info_table =  table.render({
        elem : '#reward_transaction_detail',
        cols : [[
            {field:'block_height', title:'块高'},
            {field:'block_hash', title:'块Hash'},
            {field:'reward_tx_hash', title:'分红交易Hash'},
            {field:'group_id', title:'验证组ID'},
            {field:'caster_id', title:'出块人ID'},
            {field:'members', title:'分红者列表'},
            {field:'reward_value', title:'每人分红金额'}
        ]],
        page : true,
        limit : 15
    });

    let reward_stat_table =  table.render({
        elem : '#reward_transaction_total_detail',
        cols : [[
            {field:'member_id', title:'轻节点ID'},
            {field:'reward_num', title:'验证次数'},
            {field:'total_reward_value', title:'分红总额'}
        ]],
        page : true,
        limit : 15
    });

    let cast_block_stat_table =  table.render({
        elem: '#cast_block_total_detail',
        cols: [[
            {field: 'caster_id', title: '重节点ID'},
            {field: 'cast_block_num', title: '出块次数'},
            {field: 'stake', title: '质押权益'}
        ]],
        page: true,
        limit: 15
    });

    $("#dashboard_update_div").click(function () {
        console.log('dashboard_update_switch click');
        if ($("#dashboard_update_switch").is(':checked')){
            dashboard_update_switch = true;
            updateDashboardUpdate()
        } else {
            dashboard_update_switch = false;
            updateDashboardUpdate()
        }
    });

    $("#change_host").click(function () {
        layer.prompt({
            formType: 0,
            value: HOST,
            title: '请输入新的host',
        }, function(value, index, elem){
            HOST = value;
            host_ele.text(HOST);
            layer.close(index);
            current_block_height = 0;
            blocks = [];
            block_table.reload({
                data: blocks
            });
        });
    });

    // 查询余额
    $(".query_btn").click(function () {
        let id = $(this).attr("id");
        let count = id.split("_")[2];
        $("#balance_message").text("");
        $("#balance_error").text("");
        let params = {
            "method": "Gtas_balance",
            "params": [$("#query_input_"+count).val()],
            "jsonrpc": "2.0",
            "id": "1"
        };
        $.ajax({
            type: 'POST',
            url: HOST,
            beforeSend: function (xhr) {
                xhr.setRequestHeader("Content-Type", "application/json");
            },
            data: JSON.stringify(params),
            success: function (rdata) {
                if (rdata.result !== undefined){
                    $("#balance_message").text(rdata.result);
                    $("#query_balance_"+count).text(rdata.result)
                }
                if (rdata.error !== undefined){
                    $("#balance_error").text(rdata.error.message);
                }
            },
        });
    });

    // 交易表单提交
    form.on('submit(t_form)', function(data){
        $("#t_message").text("");
        $("#t_error").text("");
        let private_key = data.field.private_key;
        let to = data.field.to;
        let value = data.field.value;
        let txdata = data.field.data;
        let t = data.field.type;
        let nonce = data.field.nonce;
        let gas = data.field.gas;
        let gas_price = data.field.gas_price;
        // if (from.length !== 42) {
        //     layer.msg("from 参数字段长度错误");
        //     return false
        // }
        // if (to.length !== 42) {
        //     layer.msg("to 参数字段长度错误");
        //     return false
        // }
        //func (api *GtasAPI) TxUnSafe(privateKey, target string, value, gas, gasprice, nonce uint64, txType int, data string) (*Result, error) {

        let params = {
            "method": "Dev_txUnSafe",
            "params": [private_key, to, parseFloat(value), parseInt(gas),parseInt(gas_price), parseInt(nonce), parseInt(t),txdata],
            "jsonrpc": "2.0",
            "id": "1"
        };

        $.ajax({
            type: 'POST',
            url: HOST,
            beforeSend: function (xhr) {
                xhr.setRequestHeader("Content-Type", "application/json");
            },
            data: JSON.stringify(params),
            success: function (rdata) {
                if (rdata.result !== undefined){
                    $("#t_message").text(rdata.result.message)
                }
                if (rdata.error !== undefined){
                    $("#t_error").text(rdata.error.message)
                }
            },
        });
        return false;
    });

    // 同步块信息
    function syncBlock(from, to) {
        if(from < 0)
            from = 0;
        if (to < 0) {
            to = 0
        }
        if (from > to) {
            return
        }
        let params = {
            "method": "Dev_getBlocks",
            "params": [from, to],
            "jsonrpc": "2.0",
            "id": "1"
        };
        $.ajax({
            type: 'POST',
            url: HOST,
            beforeSend: function (xhr) {
                xhr.setRequestHeader("Content-Type", "application/json");
            },
            async:false,
            data: JSON.stringify(params),
            success: function (rdata) {
                if (rdata.result != undefined && rdata.result != null) {
                    last_sync_block = to
                }
                if (rdata.result !== undefined){
                    retarr = rdata.result;
                    for(i = 0; i < retarr.length;i++) {
                        blocks.push(retarr[i]);
                        if (blocks.length > 100) {
                            blocks.shift()
                        }
                    }
                    block_table.reload({
                            data: blocks
                        }
                    )
                }
            },
        });
    }

    // 同步组信息
    function syncGroup(height) {
        let params = {
            "method": "Dev_getGroupsAfter",
            "params": [height],
            "jsonrpc": "2.0",
            "id": "1"
        };
        $.ajax({
            type: 'POST',
            url: HOST,
            beforeSend: function (xhr) {
                xhr.setRequestHeader("Content-Type", "application/json");
            },
            data: JSON.stringify(params),
            async: false,
            success: function (rdata) {
                if (rdata.result !== undefined && rdata.result != null){
                    retArr = rdata.result;
                    for(i = 0; i < retArr.length; i++) {
                        if (!groupIds.has(retArr[i]["id"])) {
                            groups.push(retArr[i]);
                            if (groups.length > 100) {
                                groups.shift()
                            }
                            groupIds.add(retArr[i]["id"])
                        }
                    }
                    group_table.reload({
                            data: groups
                        }
                    )
                }
                if (rdata.error !== undefined){
                    // $("#t_error").text(rdata.error.message)
                }
            },
            error: function (err) {
                console.log(err)
            }
        });
    }

    $("#query_wg_btn").click(function () {
        var h = $("#query_wg_input").val();
        if (h == null || h == undefined || h == '') {
            alert("请输入查询高度");
            return
        }
        queryWorkGroup(parseInt(h))
    });
    //查询工作组
    function queryWorkGroup(height) {
        let params = {
            "method": "Dev_getWorkGroup",
            "params": [height],
            "jsonrpc": "2.0",
            "id": "1"
        };
        $.ajax({
            type: 'POST',
            url: HOST,
            beforeSend: function (xhr) {
                xhr.setRequestHeader("Content-Type", "application/json");
            },
            data: JSON.stringify(params),
            success: function (rdata) {
                if (rdata.result !== undefined && rdata.result != null){
                    retArr = rdata.result;
                    work_group_table.reload({
                            data: retArr
                        }
                    )
                }
                if (rdata.error !== undefined){
                    // $("#t_error").text(rdata.error.message)
                }
            },
            error: function (err) {
                console.log(err)
            }
        });
    }

    $(document).on("click", "a[name='miner_oper_a']", function () {
        m = $(this).attr("method");
        t = $(this).attr("mtype");
        let params = {
            "method": m,
            "params": [parseInt(t)],
            "jsonrpc": "2.0",
            "id": "1"
        };
        text = $(this).text();
        $.ajax({
            type: 'POST',
            url: HOST,
            beforeSend: function (xhr) {
                xhr.setRequestHeader("Content-Type", "application/json");
            },
            data: JSON.stringify(params),
            success: function (rdata) {
                if (rdata.result !== undefined && rdata.result != null){
                    $(this).text("已申请"+text)
                } else {
                    alert(rdata.result)
                }
            },
            error: function (err) {
                console.log(err)
            }
        });
    });

    $("#apply_a").click(function () {
        f = $("#apply_miner_div");
        if(f.is(":hidden")){
            $(this).text("取消申请");
            f.show()
        }else{
            $(this).text("申请成为矿工");
            f.hide()
        }

    });

    $("#submit_apply").click(function () {
        stake = parseInt($("#text_stake").val());
        t = parseInt($("input[name='app_type_rd']:checked").val());
        $("#submit_result").text("");
        let params = {
            "method": "Dev_minerApply",
            "params": [stake, t],
            "jsonrpc": "2.0",
            "id": "1"
        };
        $.ajax({
            type: 'POST',
            url: HOST,
            beforeSend: function (xhr) {
                xhr.setRequestHeader("Content-Type", "application/json");
            },
            data: JSON.stringify(params),
            success: function (rdata) {
                $("#submit_result").text(rdata.result)
            },
            error: function (err) {
                console.log(err)
            }
        });
    });

    function reloadBlocksTable() {
        if (last_sync_block+1 <= current_block_height) {
            syncBlock(last_sync_block+1, current_block_height)
        }
    }
    function reloadGroupsTable() {
        if (last_sync_group+1 <= current_group_height) {
            syncGroup(last_sync_group+1)
        }
    }

    function dashboardLoad() {
        let params = {
            "method": "Dev_dashboard",
            "params": [],
            "jsonrpc": "2.0",
            "id": "1"
        };
        $.ajax({
            type: 'POST',
            url: HOST,
            beforeSend: function (xhr) {
                xhr.setRequestHeader("Content-Type", "application/json");
            },
            data: JSON.stringify(params),
            success: function (rdata) {
                d = rdata.result;
                //块高
                $("#block_height").text(d.block_height);

                clear = $("#tb_node_status").text() == "已停止" && d.node_info.status == "运行中";
                if (clear) {
                    current_block_height = 0;
                    blocks = [];
                    block_table.reload({
                        data: blocks
                    })
                }

                // for(let i=current_block_height+1; i<=d.block_height; i++) {
                //     syncBlock(i);
                // }
                current_block_height = d.block_height;

                //组高
                $("#group_height").text(d.group_height);
                current_group_height = d.group_height;

                if (clear) {
                    groups = [];
                    groupIds.clear();
                    group_table.reload({
                        data: groups
                    })
                } else {
                    // if (groups.length == 0) {
                    //     syncGroup(0)
                    // } else if (groups[groups.length-1]["height"]+1 < d.group_height) {
                    //     syncGroup(groups[groups.length-1]["height"]+1)
                    // }
                }
                //工作组数量
                $("#work_group_num").text(d.work_g_num);
                if ($("#tb_node_id").text() != d.node_info.id) {
                    //节点和质押信息
                    $("#tb_node_id").text(d.node_info.id);
                    $("#tx_send_from").val(d.node_info.id)
                }
                $("#tb_node_balance").text(d.node_info.balance);
                $("#tb_node_status").text(d.node_info.status);
                $("#tb_node_type").text(d.node_info.n_type);
                $("#tb_node_wg").text(d.node_info.w_group_num);
                $("#tb_node_ag").text(d.node_info.a_group_num);
                $("#tb_node_txnum").text(d.node_info.tx_pool_num);
                $("#tb_stake_body").html("");
                $.each(d.node_info.mort_gages, function (i, v) {
                    tr = "<tr>";
                    tr += "<td>" + v.stake + "</td>";
                    tr += "<td>" + v.type + "</td>";
                    tr += "<td>" + v.apply_height + "</td>";
                    tr += "<td>" + v.abort_height + "</td>";
                    mtype = 0;
                    if (v.type == "重节点") {
                        mtype = 1
                    }
                    if (v.abort_height > 0) {
                        tr += "<td><a href=\"javascript:void(0);\" name='miner_oper_a' method='Dev_minerRefund' mtype=" + mtype + ">退款</a></td>"
                    } else {
                        tr += "<td><a href=\"javascript:void(0);\" name='miner_oper_a' method='Dev_minerAbort' mtype=" + mtype + ">取消</a></td>"
                    }
                    tr += "</tr>";
                    $("#tb_stake_body").append(tr)
                });

                //链接节点
                let nodes_table = $("#nodes_table");
                nodes_table.empty();
                d.conns.sort(function (a, b) {
                    return parseInt(a.ip.split(".")[3]) - parseInt(b.ip.split(".")[3])
                });
                $.each(d.conns, function (i,val) {
                    nodes_table.append(
                        " <tr><td>id</td><td>ip</td><td>port</td></tr>".replace("ip", val.ip).replace("id", val.id).replace("port", val.tcp_port)
                    )
                })
            },
            error: function (err) {
                console.log(err);
                $("#tb_node_status").text("已停止");
                $("#trans_table").empty();
                $("#nodes_table").empty();
            }
        });
    }

    $("#block_height_btn").click(function () {
        var h = $("#block_height_input").val();
        if (h == null || h == undefined || h == '') {
            alert("请输入查询高度");
            return
        }
        doConsensusStat(parseInt(h))
    });

    function doConsensusStat(height) {
        let params = {
            "method": "Dev_castBlockAndRewardStat",
            "params": [height],
            "jsonrpc": "2.0",
            "id": "1"
        };
        $.ajax({
            type: 'POST',
            url: HOST,
            beforeSend: function (xhr) {
                xhr.setRequestHeader("Content-Type", "application/json");
            },
            data: JSON.stringify(params),
            success: function (rdata) {
                if (rdata.result !== undefined && rdata.result != null) {
                    renderRewardInfo(rdata.result)
                }
            },
            error: function (err) {
                console.log(err)
            }
        });
    }

    function renderRewardInfo(data) {
        var rewardes = [];
        rewardes.push(data.reward_info_at_height);

        reward_info_table.reload({
            data: rewardes
        });

        var rewardStat = [];
        for (i=0,len=data.rewardes.length; i < len; i++) {
            rewardStat.push(data.rewardes[i]);
        }
        reward_stat_table.reload({
            data: rewardStat
        });

        var castBlockStat = [];
        for (i=0,len=data.cast_blocks.length; i < len; i++) {
            castBlockStat.push(data.cast_blocks[i]);
        }
        cast_block_stat_table.reload({
            data:castBlockStat
        });
    }

    // dashboard同步数据
    // syncNodes();
    // syncTrans();
    // syncBlockHeight();
    // syncGroupHeight();
    // syncGroup(0)
    dashboardLoad();

    dashboard = setInterval(function(){
        dashboardLoad();
    },2000);

    function updateDashboardUpdate(){
        if (dashboard_update_switch){
            dashboard = setInterval(function(){
                dashboardLoad();
            },2000);
        } else{
            clearInterval(dashboard)
        }
    }

    function syncBlockLater() {
        begin = 0;
        if (blocks.length == 0) {
            syncBlock(current_block_height-20, current_block_height)
        } else if (current_block_height - last_sync_block > 100) {
            syncBlock(current_block_height-100, current_block_height)
        } else {
            reloadBlocksTable()
        }
    }

    function syncGroupLater() {
        begin = 0;
        if (groups.length == 0) {
            if (current_group_height > 100) {
                begin = current_group_height-100
            }
            syncGroup(begin)
        } else {
            reloadGroupsTable()
        }
    }

    blocktable_inter = 0;
    grouptable_inter = 0;

    element.on('tab(demo)', function(data){
        // if(data.index === 0) {
        //     clearInterval(dashboard)
        //     dashboard = setInterval(function(){
        //         dashboardLoad()
        //     },1000);
        // } else {
        //     clearInterval(dashboard)
        // }

        if (data.index == 5) {
            setTimeout(syncBlockLater, 10);
            blocktable_inter = setInterval(reloadBlocksTable, 2000);
        } else {
            clearInterval(blocktable_inter)
        }
        if (data.index == 6) {
            setTimeout(syncGroupLater, 10);
            grouptable_inter = setInterval(reloadGroupsTable, 10000);
        } else {
            clearInterval(grouptable_inter)
        }

    });

});