<%@ WebHandler Language="C#"Class="Handler" %>

using System;
using System.Web;
public class Handler :IHttpHandler {
    public void ProcessRequest (HttpContext context) {
        Object obj = context.Request.Params["action"];
		string as_action = "query";
		if(obj!=null){
			as_action = obj.ToString();
		}
		obj = context.Request.Params["us"];
		if(obj==null){
			context.Response.Write("用户名不能为空");
			return;
		}
		string as_name = obj.ToString();
        context.Response.ContentType= "text/html";
		if(as_action=="query")
		{
        	context.Response.Write((DbHelperOleDb.GetSingle(string.Format("select count(*) from tb_user where username='{0}'", as_name))).ToString());
		}
		else
		{
			context.Response.Write(AddUser(as_name));
		}
        
    }

    public bool IsReusable {
        get {
            return false;
        }
    }


    private string CheckUser(string as_name)
    {
        return (DbHelperOleDb.GetSingle(string.Format("select count(*) from tb_user where username='{0}'", as_name))).ToString();
    }

    private string AddUser(string as_name)
    {
        if (Convert.ToInt32(DbHelperOleDb.GetSingle("select count(*) from tb_user")) > 1000)
        {
            throw new ApplicationException("插入的用户记录数已经超出猫冬设置的最大笔数(1000)，请到QQ群里通知猫冬清除记录");
        }
        if (CheckUser(as_name) == "0")
        {
            return Insert(as_name);
        }
        else
        {
            return "0";
        }

    }

    private string Insert(string as_name)
    {
        try
        {
            return DbHelperOleDb.ExecuteSql(string.Format("insert into tb_user(username) values('{0}')", as_name)).ToString();
        }
        catch (Exception ex)
        {
            return "0";
        }
    }
}
