package com.moss.sdk;

import java.io.IOException;
import java.security.NoSuchAlgorithmException;
import java.security.NoSuchProviderException;
import java.security.spec.InvalidKeySpecException;
import java.util.LinkedList;
import java.util.Map;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.TimeoutException;

import com.larry.fabric.ChaincodeManager;
import com.larry.fabric.util.FabricManager;
import org.hyperledger.fabric.sdk.exception.CryptoException;
import org.hyperledger.fabric.sdk.exception.InvalidArgumentException;
import org.hyperledger.fabric.sdk.exception.ProposalException;
import org.hyperledger.fabric.sdk.exception.TransactionException;
import org.json.JSONException;
import org.json.JSONObject;

public class TestServiceImpl {
	public String chaincode(JSONObject json) {
		String fcn = null;
		try {
			fcn = json.getString("fcn");
		} catch (JSONException e) {
			e.printStackTrace();
		}
		JSONObject argJson = null;
		try {
			argJson = json.getJSONObject("arg");
		} catch (JSONException e) {
			e.printStackTrace();
		}
		Map<String, String> resultMap;
		LinkedList<String> args = new LinkedList<String>();
		String execCode = "";
		String execResult = "";
		try {
			ChaincodeManager manager = FabricManager.obtain().getManager();
			switch (fcn) {
				case "invoke":
					args.add(argJson.has("A") ? argJson.getString("A") : "");
					args.add(argJson.has("A") ? argJson.getString("A") : "");
					args.add(argJson.has("Val") ? argJson.getString("A") : "");
					String[] arguments = new  String[args.size()];
					args.toArray(arguments);
					resultMap = manager.invoke(fcn,arguments);
					break;
				case  "query":
					args.add(argJson.has("A") ? argJson.getString("A") : "");
					arguments = new String[args.size()];
					args.toArray(arguments);
					resultMap  = manager.query(fcn,arguments);
					break;



				default:
					return "No func found, check and try again";

			}
			execCode = resultMap.get("code");
			execResult = resultMap.get("data");
			if(execCode.equals("error")) {
				return execResult;


			}else {
				return  execResult;
			}
		} catch (NoSuchProviderException e) {
			e.printStackTrace();
		} catch (NoSuchAlgorithmException e) {
			e.printStackTrace();
		} catch (ExecutionException e) {
			e.printStackTrace();
		} catch (InvalidArgumentException e) {
			e.printStackTrace();
		} catch (InvalidKeySpecException e) {
			e.printStackTrace();
		} catch (JSONException e) {
			e.printStackTrace();
		} catch (CryptoException e) {
			e.printStackTrace();
		} catch (TransactionException e) {
			e.printStackTrace();
		} catch (TimeoutException e) {
			e.printStackTrace();
		} catch (ProposalException e) {
			e.printStackTrace();
		} catch (InterruptedException e) {
			e.printStackTrace();
		} catch (IOException e) {
			e.printStackTrace();
		//	return  responseFail("No func found, check and try again");

		}
		return  "No func found, check and try again";


	}
}
