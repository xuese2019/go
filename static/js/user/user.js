let pageNum = 1
let pageSize = 5
page(1)

function page(i){
    pageNum = i
    $.ajax({
        url:"/user/page/"+pageNum+"/"+pageSize,
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