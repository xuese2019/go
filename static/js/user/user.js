var pageNow
var pageSize = 5
page(1)

function page(i){
    pageNow = i
    $.ajax({
        url:"/user/page/"+i+"/"+pageSize,
        dataType:"json",
        async:true,
        data:$("#form-search-user").serialize(),
        type:"POST",
        beforeSend:function(){
            //请求前的处理
            $("#table-data").find('tr').remove()
        },
        success:function(req){
            //请求成功时处理
            $(req.data).each(function(i,e){
                $("#table-data").append("<tr><td>"+(i+1)+"</td><td>"+e.account+"</td><td></td></tr>")
            })
        },
        complete:function(){
            //请求完成的处理
        },
        error:function(err){
            //请求出错处理
            console.log(err)
        }
    });
}