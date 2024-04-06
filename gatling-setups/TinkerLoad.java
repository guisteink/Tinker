package tinker;

import io.gatling.javaapi.core.*;
import io.gatling.javaapi.http.*;

import java.time.Duration;

import static io.gatling.javaapi.core.CoreDsl.*;
import static io.gatling.javaapi.http.HttpDsl.*;

public class TinkerLoad extends Simulation {

  HttpProtocolBuilder httpProtocol = http
      .baseUrl("http://localhost:3000")
      .acceptHeader("application/json")
      .contentTypeHeader("application/json");

  // Scenario
  ScenarioBuilder scenario = scenario("Health Check")
      .exec(http("Health Check")
          .get("/health")
          .check(status().is(200)));

  // Configure the user load
  {
    setUp(
        scenario.injectOpen(
            constantUsersPerSec(100).during(Duration.ofMinutes(3)) // 100 usuários por segundo durante 3 minutos
        )).protocols(httpProtocol);
  }
}
