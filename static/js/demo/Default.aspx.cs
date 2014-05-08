using System;
using System.Data;
using System.Configuration;
using System.Web;
using System.Web.Security;
using System.Web.UI;
using System.Web.UI.WebControls;
using System.Web.UI.WebControls.WebParts;
using System.Web.UI.HtmlControls;
using System.Web.Services;

public partial class _Default : System.Web.UI.Page 
{
    protected void Page_Load(object sender, EventArgs e)
    {
        if (!IsPostBack)
        {
            Object obj = Request["us"];
            if (obj == null) return;
            us.Text = obj.ToString();
        }
    }
    protected void btn1_Click(object sender, EventArgs e)
    {
        if (CheckUser(us.Text.Trim()) == "0")
        {
            if (Insert(us.Text.Trim()) == "0")
            {
                ClientScript.RegisterStartupScript(this.GetType(), "x", "alert('ERROR');", true);
            }
            else
            {
                ClientScript.RegisterStartupScript(this.GetType(), "x", "alert('用户名插入成功');", true);
            }
        }
        else
        {
            ClientScript.RegisterStartupScript(this.GetType(), "x", "alert('用户名已存在');", true);
        }

    }

    public static string AddUser(string as_name)
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

    private static string Insert(string as_name)
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

    private static string CheckUser(string as_name)
    {
        return (DbHelperOleDb.GetSingle(string.Format("select count(*) from tb_user where username='{0}'", as_name))).ToString();
    }
}
