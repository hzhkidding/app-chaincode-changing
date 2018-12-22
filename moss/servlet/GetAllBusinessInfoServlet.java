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

public class GetAllBusinessInfoServlet extends HttpServlet {

	private  ChaincodeManager manager;


	/**
	 * 
	 */
	private static final long serialVersionUID = 1L;

	/**
	 */
	@Override
	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
		System.out.println("ceshidian3");
		response.setContentType("text/html;charset=UTF-8");
		PrintWriter out = response.getWriter();


		try {
			manager = FabricManager.obtain().getManager();
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
		String  fcn =  "getAllBusiness";
		String[] arguments = new String[1];
		try {
			Map<String,String> resultMap = manager.invoke(fcn,arguments);
			String jsonstr = resultMap.get("data");
            out.print(jsonstr);
		} catch (InvalidArgumentException e) {
			e.printStackTrace();
		} catch (ProposalException e) {
			e.printStackTrace();
		} catch (InterruptedException e) {
			e.printStackTrace();
		} catch (ExecutionException e) {
			e.printStackTrace();
		} catch (TimeoutException e) {
			e.printStackTrace();
		} catch (NoSuchAlgorithmException e) {
			e.printStackTrace();
		} catch (NoSuchProviderException e) {
			e.printStackTrace();
		} catch (InvalidKeySpecException e) {
			e.printStackTrace();
		} catch (CryptoException e) {
			e.printStackTrace();
		} catch (TransactionException e) {
			e.printStackTrace();
		}



		/*BusinessDao dao = new BusinessDao();

		List<Vector> infoList = dao.getAllBusinessInfo();
		
		List<HashMap<String,Object>> list = new ArrayList<HashMap<String,Object>>();
		for(Vector v:infoList){
			HashMap<String,Object> map = new HashMap<String, Object>();
			map.put("businessName", v.get(0));
			map.put("sendOutPrice", v.get(1));
			map.put("distributionPrice", v.get(2));
			map.put("shopHours", v.get(3));
			map.put("businessAddress", v.get(4));
			map.put("businessDepict", v.get(5));
			map.put("notice", v.get(6));
			map.put("businessScenery", v.get(7));
			map.put("logo", v.get(8));
			list.add(map);
		}
		*/
		JSONObject jsonObject = new JSONObject();
		/*try{
			JSONObject jsonObject2 = new JSONObject();
			jsonObject2.put("totalNum", list.size());
			jsonObject2.put("list", list);
			
			jsonObject.put("ret", 0);
			jsonObject.put("data", jsonObject2);
		}catch(Exception e){
			try {
				jsonObject.put("ret", 1);
			} catch (JSONException e1) {
				e1.printStackTrace();
			}
		}*/

		
	//	out.print(jsonObject);
		out.flush();
		out.close();
	}

	/**
	 */
	@Override
	public void doPost(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
		System.out.println("ddddddddd");
		throw new NotImplementedException();
	}

}
