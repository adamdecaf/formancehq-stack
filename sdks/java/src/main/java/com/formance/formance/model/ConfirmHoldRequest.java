/*
 * Formance Stack API
 * Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 
 *
 * The version of the OpenAPI document: develop
 * Contact: support@formance.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package com.formance.formance.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import java.io.IOException;

/**
 * ConfirmHoldRequest
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class ConfirmHoldRequest {
  public static final String SERIALIZED_NAME_AMOUNT = "amount";
  @SerializedName(SERIALIZED_NAME_AMOUNT)
  private Long amount;

  public static final String SERIALIZED_NAME_FINAL = "final";
  @SerializedName(SERIALIZED_NAME_FINAL)
  private Boolean _final;

  public ConfirmHoldRequest() {
  }

  public ConfirmHoldRequest amount(Long amount) {
    
    this.amount = amount;
    return this;
  }

   /**
   * Define the amount to transfer.
   * @return amount
  **/
  @javax.annotation.Nullable

  public Long getAmount() {
    return amount;
  }


  public void setAmount(Long amount) {
    this.amount = amount;
  }


  public ConfirmHoldRequest _final(Boolean _final) {
    
    this._final = _final;
    return this;
  }

   /**
   * Define a final confirmation. Remaining funds will be returned to the wallet.
   * @return _final
  **/
  @javax.annotation.Nullable

  public Boolean getFinal() {
    return _final;
  }


  public void setFinal(Boolean _final) {
    this._final = _final;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ConfirmHoldRequest confirmHoldRequest = (ConfirmHoldRequest) o;
    return Objects.equals(this.amount, confirmHoldRequest.amount) &&
        Objects.equals(this._final, confirmHoldRequest._final);
  }

  @Override
  public int hashCode() {
    return Objects.hash(amount, _final);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ConfirmHoldRequest {\n");
    sb.append("    amount: ").append(toIndentedString(amount)).append("\n");
    sb.append("    _final: ").append(toIndentedString(_final)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}

