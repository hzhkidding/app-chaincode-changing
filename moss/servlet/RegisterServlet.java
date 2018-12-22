package com.moss.servlet;

import com.larry.fabric.ChaincodeManager;
import com.larry.fabric.util.FabricManager;
import org.hyperledger.fabric.sdk.exception.CryptoException;
import org.hyperledger.fabric.sdk.exception.InvalidArgumentException;
import org.hyperledger.fabric.sdk.exception.ProposalException;
import org.hyperledger.fabric.sdk.exception.TransactionException;
import org.json.JSONObject;
import sun.reflect.generics.reflectiveObjects.NotImplementedException;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;
import java.io.PrintWriter;
import java.security.NoSuchAlgorithmException;
import java.security.NoSuchProviderException;
import java.security.spec.InvalidKeySpecException;
import java.util.Map;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.TimeoutException;
/**
 * 
 * 项目名称：MealOrderSystemServer    
 * 类名称：RegisterServlet    
 * 类描述：注册的控制类    
 * 创建人：Nicky
 * 创建时间：2016年7月4日 下午1:40:15      
 * @version
 */
public class RegisterServlet extends HttpServlet {

	/**
	 * 
	 */
	public static ChaincodeManager manager;
	private static final long serialVersionUID = 1L;


	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
		throw new NotImplementedException();
	}

	
	public void doPost(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
     System.out.println("hahahanishishabifasf");
		response.setContentType("text/html;charset=UTF-8");
		String account = request.getParameter("account");
		String password = request.getParameter("password");
		System.out.println(account);
		String rank = "LV1";
		String credit = "0";
		String phone = "哈哈哈";
		String imgPath = "/admin/images/face.jpg";
		
		JSONObject jObject = new JSONObject();

      try {
             manager = FabricManager.obtain().getManager();
            String[] arguments = new String[6];
            arguments[0] = account;
            arguments[1] = password;
            arguments[2] = rank;
            arguments[3] = String.valueOf(credit);
            arguments[4] = phone;
            arguments[5] = imgPath;
            String fcn = "register";
			try {
				//Map<String,String> resultMap = manager.invoke(fcn,arguments);
				Map<String, String> resultMap = manager.query(fcn, arguments);
				System.out.println(resultMap.get("data"));
			} catch (ProposalException e) {
				e.printStackTrace();
			}
	  } catch (CryptoException e) {
            e.printStackTrace();
        } catch (InvalidArgumentException e) {
            e.printStackTrace();
        } catch (NoSuchAlgorithmException e) {
            e.printStackTrace();
        } catch (NoSuchProviderException e) {
            e.printStackTrace();
        } catch (InvalidKeySpecException e) {
            e.printStackTrace();
        } catch (TransactionException e) {
            e.printStackTrace();
        }
       /*try {
			
			MemberDao dao = new MemberDao();
			Member member = new Member();
			member.setMemberID("1");
			member.setPassword(password);
			member.setRank(rank);
			member.setCredit(credit);
			member.setImgPath(imgPath);
			
			dao.register(member);
			
			jObject.put("ret", 0);
			jObject.put("msg", "ok");
			jObject.put("data", "");
			
		} catch (Exception e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
			try
			{
				jObject.put("ret", 1);
				jObject.put("msg", "error");
				jObject.put("data", "");
			} catch (JSONException ex)
			{
				ex.printStackTrace();
			}
		}
*/
		PrintWriter out = response.getWriter();

		//out.println(jObject);
		out.flush();
		out.close();
	}
}
