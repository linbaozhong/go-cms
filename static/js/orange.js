
$(function(){
	var ue = UM.getEditor('form-content');

	$('#defaultForm').isHappy({
          fields: {
            '#form-title': {
            	required: true,
            	message: '标题是必须填写的'
            },
            '#form-date': {
            	required: true,
            	message: '日期格式不对',
            	test: happy.date
           	},
           	'#form-content': {
           		required: true,
           		message: '内容是必须填写的'
           	}
       	}
    });
    $('#defaultForm').on("submit", function(){
    	var $this = $(this),
    		$data = $this.serialize()
    	$.post("/admin/article/create", $data, function(data){
    		//console.log(data.Ok,data.Key,data.Data)
    		if( data.Ok ){
    			alert("提交成功!")
    			window.location.href='/admin/article';
    		}else{
    			console.log(data.Key + "：" + data.Data)
    		}
    	})
    	.complete(function(){})
    })
});
